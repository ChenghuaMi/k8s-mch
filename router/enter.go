/**
 * @author mch
 */

package router

import (
	"k8s-mch/router/example"
	"k8s-mch/router/k8s"
)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8sRouterGroup k8s.K8sRouter
}

var RouterGroupApp = new(RouterGroup)