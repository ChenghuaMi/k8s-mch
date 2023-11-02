/**
 * @author mch
 */

package response

import (
	"k8s-mch/model/base"
	corev1 "k8s.io/api/core/v1"
)

type Node struct {
	Name string `json:"name"`
	Status string `json:"status"`
	Age int64 `json:"age"`
	InternalIp string `json:"internalIp"`
	ExternalIp string `json:"externalIp"`

	Version string `json:"version"`
	OsImage string `json:"osImage"`
	KernelVersion string `json:"kernelVersion"`
	ContainerRuntime string `json:"containerRuntime"`
	Labels  []base.ListMapItem `json:"labels"`
	Taints  []corev1.Taint `json:"taints"`
}
