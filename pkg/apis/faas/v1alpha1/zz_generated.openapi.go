// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarService":       schema_pkg_apis_faas_v1alpha1_JarService(ref),
		"github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceSpec":   schema_pkg_apis_faas_v1alpha1_JarServiceSpec(ref),
		"github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceStatus": schema_pkg_apis_faas_v1alpha1_JarServiceStatus(ref),
	}
}

func schema_pkg_apis_faas_v1alpha1_JarService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JarService is the Schema for the jarservices API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceSpec", "github.com/slinkydeveloper/knative-jar-operator/pkg/apis/faas/v1alpha1.JarServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_faas_v1alpha1_JarServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JarServiceSpec defines the desired state of JarService",
				Properties: map[string]spec.Schema{
					"jarLocation": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"jarLocation"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_faas_v1alpha1_JarServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JarServiceStatus defines the observed state of JarService",
				Properties: map[string]spec.Schema{
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"nodes"},
			},
		},
		Dependencies: []string{},
	}
}
