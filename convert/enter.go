/**
 * @author mch
 */

package convert

import (
	"k8s-mch/convert/node"
	"k8s-mch/convert/pod"
)

type ConvertGroup struct {
	PodConvert pod.PodConvert
	NodeConvert  node.NodeConvert
}

var ConvertGroupApp = new(ConvertGroup)