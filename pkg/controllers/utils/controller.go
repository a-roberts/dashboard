package utils

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
	"github.com/tektoncd/dashboard/pkg/endpoints"
	"github.com/tektoncd/dashboard/pkg/logging"
<<<<<<< HEAD
	"github.com/tektoncd/dashboard/pkg/utils"
=======
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

<<<<<<< HEAD
func NewController(kind string, informer cache.SharedIndexInformer, onCreated, onUpdated, onDeleted broadcaster.MessageType, filter func(interface{}, bool) interface{}) {
	logging.Log.Debug("In NewController")

	if filter == nil {
		filter = func(obj interface{}, skipDeletedCheck bool) interface{} {
			return obj
		}
	}

=======
func NewController(kind string, informer cache.SharedIndexInformer, onCreated, onUpdated, onDeleted broadcaster.MessageType) {
	logging.Log.Debug("In NewController")

>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			logging.Log.Debugf("Controller detected %s '%s' created", kind, obj.(metav1.Object).GetName())
			data := broadcaster.SocketData{
				MessageType: onCreated,
<<<<<<< HEAD
				Payload:     filter(obj, true),
=======
				Payload:     obj,
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
			}
			endpoints.ResourcesChannel <- data
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldSecret, newSecret := oldObj.(metav1.Object), newObj.(metav1.Object)
			// If resourceVersion differs between old and new, an actual update event was observed
			if oldSecret.GetResourceVersion() != newSecret.GetResourceVersion() {
				logging.Log.Debugf("Controller detected %s '%s' updated", kind, oldSecret.GetName())
				data := broadcaster.SocketData{
					MessageType: onUpdated,
<<<<<<< HEAD
					Payload:     filter(newObj, true),
=======
					Payload:     newObj,
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
				}
				endpoints.ResourcesChannel <- data
			}
		},
		DeleteFunc: func(obj interface{}) {
<<<<<<< HEAD
			logging.Log.Debugf("Controller detected %s '%s' deleted", kind, utils.GetDeletedObjectMeta(obj).GetName())
			data := broadcaster.SocketData{
				MessageType: onDeleted,
				Payload:     filter(obj, false),
=======
			logging.Log.Debugf("Controller detected %s '%s' deleted", kind, GetDeletedObjectMeta(obj).GetName())
			data := broadcaster.SocketData{
				MessageType: onDeleted,
				Payload:     obj,
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
			}
			endpoints.ResourcesChannel <- data
		},
	})
}
