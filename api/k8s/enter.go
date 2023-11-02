/**
 * @author mch
 */

package k8s

import (
	"k8s-mch/service"
	"k8s-mch/validate"
)

type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
}
var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.PodServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.PodServiceGroupApp.NodeServiceGroup.NodeService

