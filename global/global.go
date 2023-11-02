/**
 * @author mch
 */

package global

import (
	"k8s-mch/config"
	"k8s.io/client-go/kubernetes"
)

var (
	Conf config.Server
	KubeConfigSet *kubernetes.Clientset
)
