/**
 * @author mch
 */

package initiallize

import (
	"github.com/gin-gonic/gin"
	"k8s-mch/router"
)

func Routers() *gin.Engine{
	r := gin.Default()
	router.RouterGroupApp.ExampleRouterGroup.InitRouter(r)
	router.RouterGroupApp.K8sRouterGroup.InitRouter(r)
	return r
}