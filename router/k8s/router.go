/**
 * @author mch
 */

package k8s

import (
	"github.com/gin-gonic/gin"
	"k8s-mch/api"
)

type K8sRouter struct {

}

func(*K8sRouter) InitRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	group.GET("/podList",api.ApiGroupApp.K8sApiGroup.GetPodList)
	group.GET("/namespaceList",api.ApiGroupApp.K8sApiGroup.GetNamespaceList)
	group.POST("/savePod",api.ApiGroupApp.K8sApiGroup.CreateOrUpdatePod)

	//node
	group.GET("/nodeList",api.ApiGroupApp.K8sApiGroup.GetNodeDetailOrList)
}
