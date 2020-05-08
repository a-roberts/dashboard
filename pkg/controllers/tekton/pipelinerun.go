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

// NewPipelineRunController registers a Tekton controller/informer for
// pipelineRuns on the sharedTektonInformerFactory
func NewPipelineRunController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewPipelineRunController")

<<<<<<< HEAD
func pipelineRunDeleted(obj interface{}) {
	logging.Log.Debugf("PipelineRun Controller detected pipelineRun '%s' deleted", obj.(*v1alpha1.PipelineRun).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.PipelineRunDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"pipeline run",
		sharedTektonInformerFactory.Tekton().V1alpha1().PipelineRuns().Informer(),
		broadcaster.PipelineRunCreated,
		broadcaster.PipelineRunUpdated,
		broadcaster.PipelineRunDeleted,
		nil,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
