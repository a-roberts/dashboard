package tekton

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
<<<<<<< HEAD
	"github.com/tektoncd/dashboard/pkg/endpoints"
=======
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
	logging "github.com/tektoncd/dashboard/pkg/logging"
	tektonresourceinformer "github.com/tektoncd/pipeline/pkg/client/resource/informers/externalversions"
)

// NewPipelineResourceController registers a Tekton controller/informer for
// pipelineResources on the sharedTektonInformerFactory
func NewPipelineResourceController(sharedTektonInformerFactory tektonresourceinformer.SharedInformerFactory) {
	logging.Log.Debug("In NewPipelineResourceController")

<<<<<<< HEAD
func pipelineResourceDeleted(obj interface{}) {
	logging.Log.Debugf("PipelineResource Controller detected pipelineResource '%s' deleted", obj.(*v1alpha1.PipelineResource).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.PipelineResourceDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"pipeline resource",
		sharedTektonInformerFactory.Tekton().V1alpha1().PipelineResources().Informer(),
		broadcaster.PipelineResourceCreated,
		broadcaster.PipelineResourceUpdated,
		broadcaster.PipelineResourceDeleted,
		nil,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
