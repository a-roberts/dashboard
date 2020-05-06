/*
Copyright 2019 The Tekton Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

// NewClusterTaskController registers a Tekton controller/informer for cluster
// tasks on the sharedTektonInformerFactory
func NewClusterTaskController(sharedTektonInformerFactory tektoninformer.SharedInformerFactory) {
	logging.Log.Debug("In NewClusterTaskController")

<<<<<<< HEAD
func clusterTaskDeleted(obj interface{}) {
	logging.Log.Debugf("Cluster Task Controller detected clusterTask '%s' deleted", obj.(*v1alpha1.ClusterTask).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.ClusterTaskDeleted,
		Payload:     obj,
	}
	endpoints.ResourcesChannel <- data
=======
	utils.NewController(
		"cluster task",
		sharedTektonInformerFactory.Tekton().V1alpha1().ClusterTasks().Informer(),
		broadcaster.ClusterTaskCreated,
		broadcaster.ClusterTaskUpdated,
		broadcaster.ClusterTaskDeleted,
	)
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
}
