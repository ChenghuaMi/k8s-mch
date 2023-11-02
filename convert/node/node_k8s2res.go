/**
 * @author mch
 */

package node

import (
	"k8s-mch/model/base"
	node_res "k8s-mch/model/node/response"
	corev1 "k8s.io/api/core/v1"
)

type NodeK8s2Res struct {

}
func(this *NodeK8s2Res) getNodeStatus(nodeConditions []corev1.NodeCondition) string {
	nodeStatus := "NotReady"
	for _,condition := range nodeConditions {
		if condition.Type == "Ready" && condition.Status == "True" {
			nodeStatus = "Ready"
			break
		}
	}
	return nodeStatus
}
func(this *NodeK8s2Res) getNodeIp(addresses []corev1.NodeAddress,addressType  corev1.NodeAddressType) string {
	nodeIp := "<none>"
	for _,item := range addresses {
		if item.Type == addressType {
			nodeIp = item.Address
		}
	}
	return nodeIp
}
func(this *NodeK8s2Res) mapToList(m map[string]string) []base.ListMapItem {
	res := make([]base.ListMapItem,0)
	for k,v := range m {
		res = append(res,base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return res
}
func(this *NodeK8s2Res) GetNodeDetail(nodeK8s corev1.Node) node_res.Node {
	nodeRes := this.GetNodeRes(nodeK8s)

	nodeRes.Taints = nodeK8s.Spec.Taints
	nodeRes.Labels = this.mapToList(nodeK8s.Labels)
	return nodeRes
}
func(this *NodeK8s2Res) GetNodeRes(nodeK8s corev1.Node) node_res.Node {
	nodeInfo := nodeK8s.Status.NodeInfo
	return node_res.Node{
		Name:             nodeK8s.Name,
		Status:           this.getNodeStatus(nodeK8s.Status.Conditions),
		Age:              nodeK8s.CreationTimestamp.Unix(),
		InternalIp:       this.getNodeIp(nodeK8s.Status.Addresses,corev1.NodeInternalIP),
		ExternalIp:       this.getNodeIp(nodeK8s.Status.Addresses,corev1.NodeExternalIP),
		Version:          nodeInfo.KubeletVersion,
		OsImage:          nodeInfo.OSImage,
		KernelVersion:    nodeInfo.KernelVersion,
		ContainerRuntime: nodeInfo.ContainerRuntimeVersion,
	}
}
