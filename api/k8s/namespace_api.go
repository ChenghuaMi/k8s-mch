/**
 * @author mch
 */

package k8s

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s-mch/global"
	namespace_res "k8s-mch/model/namespace/response"
	"k8s-mch/response"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceApi struct {

}
func(*NamespaceApi) GetNamespaceList(c *gin.Context) {
	list,err := global.KubeConfigSet.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	namespaceList := make([]namespace_res.Namespace,0)
	for _,item := range list.Items {
		namespaceList = append(namespaceList,namespace_res.Namespace{
			Name:            item.Name,
			CreationTimestamp: item.CreationTimestamp.Unix(),
			Status:          string(item.Status.Phase),
		})
	}
	response.SuccessWithDetailed(c,"ok",namespaceList)
}
