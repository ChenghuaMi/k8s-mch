/**
 * @author mch
 */

package validate

import (
	"errors"
	pod_req "k8s-mch/model/pod/request"
)

type PodValidate struct {

}
func(*PodValidate) Validate(podReq *pod_req.Pod) error {
	if podReq.Base.Name == "" {
		return errors.New("pod name 不能为空")
	}
	if len(podReq.Containers) == 0 {
		return errors.New("请定义pod 容器信息")
	}
	if len(podReq.InitContainers) > 0 {
		for index,container := range podReq.InitContainers {
			if container.Name == "" {
				return errors.New("initContainer 没有定义容器名称")
			}
			if container.Image == "" {
				return errors.New("initContainer 没有定镜像")
			}
			if container.ImagePullPolicy == "" {
				podReq.InitContainers[index].ImagePullPolicy = "IfNotPresent"
			}
		}
	}
	if len(podReq.Containers) > 0 {
		for index,container := range podReq.Containers {
			if container.Name == "" {
				return errors.New("container 没有定义容器名称")
			}
			if container.Image == "" {
				return errors.New("container 没有定镜像")
			}
			if container.ImagePullPolicy == "" {
				podReq.Containers[index].ImagePullPolicy = "IfNotPresent"
			}
		}
	}
	if podReq.Base.RestartPolicy == "" {
		podReq.Base.RestartPolicy = "Always"
	}
	return nil
}