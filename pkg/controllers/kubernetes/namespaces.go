package kubernetes

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
<<<<<<< HEAD
	"github.com/tektoncd/dashboard/pkg/endpoints"
	logging "github.com/tektoncd/dashboard/pkg/logging"
	v1 "k8s.io/api/core/v1"
=======
	"github.com/tektoncd/dashboard/pkg/controllers/utils"
	"github.com/tektoncd/dashboard/pkg/logging"
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
	k8sinformer "k8s.io/client-go/informers"
)

// NewNamespaceController registers the K8s shared informer that reacts to
// create and delete events for namespaces
func NewNamespaceController(sharedK8sInformerFactory k8sinformer.SharedInformerFactory) {
	logging.Log.Debug("In NewNamespaceController")
<<<<<<< HEAD
	k8sInformer := sharedK8sInformerFactory.Core().V1().Namespaces()
	k8sInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    namespaceCreated,
		DeleteFunc: namespaceDeleted,
	})
}

func namespaceCreated(obj interface{}) {
	logging.Log.Debugf("Namespace Controller detected namespace '%s' created", obj.(*v1.Namespace).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.NamespaceCreated,
		Payload:     obj,
	}

	endpoints.ResourcesChannel <- data
}

func namespaceDeleted(obj interface{}) {
	logging.Log.Debugf("Namespace Controller detected namespace '%s' deleted", obj.(*v1.Namespace).Name)
	data := broadcaster.SocketData{
		MessageType: broadcaster.NamespaceDeleted,
		Payload:     obj,
	}
=======
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication

	utils.NewController(
		"namespace",
		sharedK8sInformerFactory.Core().V1().Namespaces().Informer(),
		broadcaster.NamespaceCreated,
		broadcaster.NamespaceUpdated,
		broadcaster.NamespaceDeleted,
		nil,
	)
}
