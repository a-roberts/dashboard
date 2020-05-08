package utils

import (
	"github.com/tektoncd/dashboard/pkg/broadcaster"
	"github.com/tektoncd/dashboard/pkg/endpoints"
	"github.com/tektoncd/dashboard/pkg/logging"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/tektoncd/dashboard/pkg/utils"
=======
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
	"github.com/tektoncd/dashboard/pkg/utils"
>>>>>>> 410c7ca... Filter resources before returning to the client
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

<<<<<<< HEAD
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
=======
func NewController(kind string, informer cache.SharedIndexInformer, onCreated, onUpdated, onDeleted broadcaster.MessageType, filter func(interface{}, bool) interface{}) {
	logging.Log.Debug("In NewController")

	if filter == nil {
		filter = func(obj interface{}, skipDeletedCheck bool) interface{} {
			return obj
		}
	}

>>>>>>> 410c7ca... Filter resources before returning to the client
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			logging.Log.Debugf("Controller detected %s '%s' created", kind, obj.(metav1.Object).GetName())
			data := broadcaster.SocketData{
				MessageType: onCreated,
<<<<<<< HEAD
<<<<<<< HEAD
				Payload:     filter(obj, true),
=======
				Payload:     obj,
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
				Payload:     filter(obj, true),
>>>>>>> 410c7ca... Filter resources before returning to the client
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
<<<<<<< HEAD
					Payload:     filter(newObj, true),
=======
					Payload:     newObj,
>>>>>>> ddd6c4f... Refactor controllers to reduce code duplication
=======
					Payload:     filter(newObj, true),
>>>>>>> 410c7ca... Filter resources before returning to the client
				}
				endpoints.ResourcesChannel <- data
			}
		},
		DeleteFunc: func(obj interface{}) {
<<<<<<< HEAD
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
=======
			logging.Log.Debugf("Controller detected %s '%s' deleted", kind, utils.GetDeletedObjectMeta(obj).GetName())
			data := broadcaster.SocketData{
				MessageType: onDeleted,
				Payload:     filter(obj, false),
>>>>>>> 410c7ca... Filter resources before returning to the client
			}
			endpoints.ResourcesChannel <- data
		},
	})
}
