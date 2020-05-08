package tekton

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
	logging "github.com/tektoncd/dashboard/pkg/logging"
	tektoninformer "github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
)

// NewTaskRunController registers a Tekton controller/informer for taskRuns on
// the sharedTektonInformerFactory
func NewTaskRunController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewTaskRunController")

	utils.NewController(
		"task run",
		sharedTektonInformerFactory.Tekton().V1alpha1().TaskRuns().Informer(),
		broadcaster.TaskRunCreated,
		broadcaster.TaskRunUpdated,
		broadcaster.TaskRunDeleted,
<<<<<<< HEAD
<<<<<<< HEAD
		nil,
=======
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
		nil,
>>>>>>> 410c7ca... Filter resources before returning to the client
	)
}
