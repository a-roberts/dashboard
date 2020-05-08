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

// NewTaskController registers a Tekton controller/informer for tasks on
// the sharedTektonInformerFactory
func NewTaskController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewTaskController")

<<<<<<< HEAD
func taskDeleted(obj interface{}) {
	logging.Log.Debugf("Task Controller detected task '%s' deleted", obj.(*v1alpha1.Task).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.TaskDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"task",
		sharedTektonInformerFactory.Tekton().V1alpha1().Tasks().Informer(),
		broadcaster.TaskCreated,
		broadcaster.TaskUpdated,
		broadcaster.TaskDeleted,
		nil,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
