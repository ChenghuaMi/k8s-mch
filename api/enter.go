/**
 * @author mch
 */

package api

import (
	"k8s-mch/api/example"
	"k8s-mch/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8sApiGroup k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)