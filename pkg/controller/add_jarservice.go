package controller

import (
	"github.com/slinkydeveloper/knative-jar-operator/pkg/controller/jarservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, jarservice.Add)
}
