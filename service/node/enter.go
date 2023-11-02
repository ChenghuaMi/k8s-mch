/**
 * @author mch
 */

package node

import "k8s-mch/convert"

type NodeServiceGroup struct {
	NodeService NodeService
}

var nodeConvert = convert.ConvertGroupApp.NodeConvert