package jarservice

import (
	"context"
	"fmt"

	knv1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
	knv1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"

	jarv1alpha1 "github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_jarservice")

// Add creates a new JarService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileJarService{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("jarservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource JarService
	err = c.Watch(&source.Kind{Type: &jarv1alpha1.JarService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner JarService
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &jarv1alpha1.JarService{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileJarService implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileJarService{}

// ReconcileJarService reconciles a JarService object
type ReconcileJarService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a JarService object and makes changes based on the state read
// and what is in the JarService.Spec
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileJarService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling JarService")

	// Fetch the JarService instance
	jarService := &jarv1alpha1.JarService{}
	err := r.client.Get(context.TODO(), request.NamespacedName, jarService)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatix	cally garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if Knative service already exists for this JarService
	found := &knv1alpha1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: jarService.Name, Namespace: jarService.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		// No service for this function exists. Create a new one
		service, err := r.newServiceForJarService(jarService)
		if err != nil {
			return reconcile.Result{}, err
		}

		reqLogger.Info("Creating a new knative Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		err = r.client.Create(context.TODO(), service)
		if err != nil {
			reqLogger.Error(err, "Failed to create new Service.", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
			return reconcile.Result{}, err
		}

		// Service created successfully - return and requeue
		return reconcile.Result{Requeue: true}, nil
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Service for JSFunction")
		return reconcile.Result{}, err
	}

	// TODO log
	reqLogger.Info("Skip reconcile: Service already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func (r *ReconcileJarService) newServiceForJarService(cr *jarv1alpha1.JarService) (*knv1alpha1.Service, error) {
	service := &knv1alpha1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: knv1alpha1.ServiceSpec{
			ConfigurationSpec: knv1alpha1.ConfigurationSpec{
				Template: &knv1alpha1.RevisionTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{"sidecar.istio.io/inject": "false"},
					},
					Spec: knv1alpha1.RevisionSpec{
						RevisionSpec: knv1beta1.RevisionSpec{
							PodSpec: corev1.PodSpec{
								Containers: []corev1.Container{{
									Image: "slinkydeveloper/faas-java-runtime-image:8",
									Name:  fmt.Sprintf("java-8-%s", cr.Name),
									Ports: []corev1.ContainerPort{{
										ContainerPort: 8080,
									}},
									Env: []corev1.EnvVar{{
										Name: "IMAGE_LINK",
										Value: cr.Spec.JarLocation,
									}},
								}},
							},
						},
					},
				},
			},
		},
	}

	// Set JarService instance as the owner and controller of service
	if err := controllerutil.SetControllerReference(cr, service, r.scheme); err != nil {
		return nil, err
	}

	return service, nil
}