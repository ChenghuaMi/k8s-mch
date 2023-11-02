/**
 * @author mch
 */

package example

import (
	"github.com/gin-gonic/gin"
	"k8s-mch/api"
)
// ExampleRouter luoyou
type ExampleRouter struct {

}
func(*ExampleRouter) InitRouter(r *gin.Engine) {
	group := r.Group("/example")
	group.GET("/pong",api.ApiGroupApp.ExampleApiGroup.ExampleTest)
}
