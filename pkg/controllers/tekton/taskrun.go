package tekton

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
<<<<<<< HEAD
	"github.com/tektoncd/dashboard/pkg/endpoints"
=======
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
	logging "github.com/tektoncd/dashboard/pkg/logging"
	tektoninformer "github.com/tektoncd/pipeline/pkg/client/informers/externalversions"
)

// NewTaskRunController registers a Tekton controller/informer for taskRuns on
// the sharedTektonInformerFactory
func NewTaskRunController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewTaskRunController")

<<<<<<< HEAD
func taskRunDeleted(obj interface{}) {
	logging.Log.Debugf("TaskRun Controller detected taskRun '%s' deleted", obj.(*v1alpha1.TaskRun).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.TaskRunDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"task run",
		sharedTektonInformerFactory.Tekton().V1alpha1().TaskRuns().Informer(),
		broadcaster.TaskRunCreated,
		broadcaster.TaskRunUpdated,
		broadcaster.TaskRunDeleted,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
