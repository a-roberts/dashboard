package kubernetes

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
<<<<<<< HEAD
<<<<<<< HEAD
	controllerUtils "github.com/tektoncd/dashboard/pkg/controllers/utils"
	"github.com/tektoncd/dashboard/pkg/logging"
	"github.com/tektoncd/dashboard/pkg/utils"
=======
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
	"github.com/tektoncd/dashboard/pkg/logging"
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
	controllerUtils "github.com/tektoncd/dashboard/pkg/controllers/utils"
	"github.com/tektoncd/dashboard/pkg/logging"
	"github.com/tektoncd/dashboard/pkg/utils"
>>>>>>> 410c7ca... Filter resources before returning to the client
	k8sinformer "k8s.io/client-go/informers"
)

// NewSecretController registers the K8s shared informer that reacts to
// create, update and delete events for secrets
func NewSecretController(sharedK8sInformerFactory k8sinformer.SharedInformerFactory) {
	logging.Log.Debug("In NewSecretController")

<<<<<<< HEAD
<<<<<<< HEAD
	controllerUtils.NewController(
=======
	utils.NewController(
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
	controllerUtils.NewController(
>>>>>>> 410c7ca... Filter resources before returning to the client
		"secret",
		sharedK8sInformerFactory.Core().V1().Secrets().Informer(),
		broadcaster.SecretCreated,
		broadcaster.SecretUpdated,
		broadcaster.SecretDeleted,
<<<<<<< HEAD
<<<<<<< HEAD
		utils.SanitizeSecret,
=======
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
		utils.SanitizeSecret,
>>>>>>> 410c7ca... Filter resources before returning to the client
	)
}
