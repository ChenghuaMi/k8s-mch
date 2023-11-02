/**
 * @author mch
 */

package service

import (
	"k8s-mch/service/node"
	"k8s-mch/service/pod"
)

type ServiceGroup struct {
	PodServiceGroup pod.PodServiceGroup
	NodeServiceGroup node.NodeServiceGroup
}

var PodServiceGroupApp = new(ServiceGroup)