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

// NewPipelineController registers a Tekton controller/informer for pipelines on
// the sharedTektonInformerFactory
func NewPipelineController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewPipelineController")

<<<<<<< HEAD
func pipelineDeleted(obj interface{}) {
	logging.Log.Debugf("Pipeline Controller detected pipeline '%s' deleted", obj.(*v1alpha1.Pipeline).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.PipelineDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"pipeline",
		sharedTektonInformerFactory.Tekton().V1alpha1().Pipelines().Informer(),
		broadcaster.PipelineCreated,
		broadcaster.PipelineUpdated,
		broadcaster.PipelineDeleted,
		nil,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
