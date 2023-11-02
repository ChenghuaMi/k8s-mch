/**
 * @author mch
 */

package pod

import (
	"context"
	"fmt"
	"k8s-mch/global"
	pod_req "k8s-mch/model/pod/request"
	pod_res "k8s-mch/model/pod/response"
	corev1 "k8s.io/api/core/v1"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"strings"
)

type PodService struct {

}
func(*PodService) GetPodList(namespace string) ([]pod_res.PodListItem,error) {
	//list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	return nil,err
	//}
	//result := make([]pod_res.PodListItem,0)
	return nil,nil
}
func(*PodService) CreateOrUpdate(podReq pod_req.Pod) (string,error) {
	k8sPod := podConvert.PodReq2K8s(podReq)
	ctx := context.TODO()
	podApi := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace)
	if k8sGetPod,err := podApi.Get(ctx,k8sPod.Name,metav1.GetOptions{});err == nil {
		// 校验 pod 参数是否合理
		k8sPodCopy := *k8sPod
		k8sPodCopy.Name = k8sPod.Name + "-validate"
		_, err := podApi.Create(ctx, &k8sPodCopy, metav1.CreateOptions{
			DryRun: []string{metav1.DryRunAll},
		})
		if err != nil {
			msg := err.Error()

			return msg,err
		}
		//delete
		err = podApi.Delete(ctx,k8sPod.Name,metav1.DeleteOptions{})
		if err != nil {
			msg := err.Error()

			return msg,err
		}
		labelSelector := make([]string,0)
		for k,v := range k8sGetPod.Labels {
			labelSelector = append(labelSelector,fmt.Sprintf("%s=%s",k,v))
		}
		//pod  terminating 状态，监听 pod 删除完毕创建 pod
		watcher, err := podApi.Watch(ctx, metav1.ListOptions{
			LabelSelector: strings.Join(labelSelector, ","),
		})
		if err != nil {
			msg := err.Error()

			return msg,err
		}
		for event := range watcher.ResultChan() {
			k8sPodChan := event.Object.(*corev1.Pod)
			// 查询k8s pod 是否删除
			_, err := podApi.Get(ctx, k8sPod.Name, metav1.GetOptions{})
			if k8sError.IsNotFound(err) {
				if createdPod,err := podApi.Create(context.TODO(),k8sPod,metav1.CreateOptions{});err != nil {
					msg := fmt.Sprintf("pod 【%s-%s】更新失败,errors %s",k8sPod.Namespace,k8sPod.Name,err.Error())
					return msg,err
				}else{
					msg := fmt.Sprintf("pod 【%s-%s】更新成功",createdPod.Namespace,createdPod.Name)
					return msg,err
				}
			}
			switch event.Type {
			case watch.Deleted:
				if k8sPodChan.Name != k8sPod.Name {
					continue
				}
				if createdPod,err := podApi.Create(context.TODO(),k8sPod,metav1.CreateOptions{});err != nil {
					msg := fmt.Sprintf("pod 【%s-%s】更新失败,errors %s",k8sPod.Namespace,k8sPod.Name,err.Error())
					return msg,err
				}else{
					msg := fmt.Sprintf("pod 【%s-%s】更新成功",createdPod.Namespace,createdPod.Name)
					return msg,err

				}
			}
		}
		//重新创建
	}else {
		if createdPod,err := podApi.Create(context.TODO(),k8sPod,metav1.CreateOptions{});err != nil {
			msg := fmt.Sprintf("pod 【%s-%s】创建失败,errors %s",k8sPod.Namespace,k8sPod.Name,err.Error())
			return msg,err

		}else{
			msg := fmt.Sprintf("pod 【%s-%s】创建成功",createdPod.Namespace,createdPod.Name)
			return msg,err

		}
	}
	return "",nil
}