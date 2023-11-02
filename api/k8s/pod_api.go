/**
 * @author mch
 */

package k8s

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s-mch/global"
	pod_req "k8s-mch/model/pod/request"
	"k8s-mch/response"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type PodApi struct {

}
func(*PodApi) CreateOrUpdatePod(c *gin.Context) {
		var podReq pod_req.Pod
		if err := c.ShouldBind(&podReq);err != nil {
			response.FailWithMessage(c,err.Error())
			return
		}
		if err := podValidate.Validate(&podReq);err != nil {
			response.FailWithMessage(c,err.Error())
			return
		}
		msg,err := podService.CreateOrUpdate(podReq)
		if err != nil {
			response.FailWithMessage(c,msg)
			return
		}
		response.SuccessWithMessage(c,msg)

}
func(*PodApi) GetPodList(c *gin.Context) {
	list,err := global.KubeConfigSet.CoreV1().Pods("").List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _,item := range list.Items {
		fmt.Println("namespace:",item.Namespace,"name:",item.Name)
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"ok",
	})
}