/**
 * @author mch
 */

package pod

import "k8s-mch/convert"

type PodServiceGroup struct {
	PodService PodService
}

var podConvert = convert.ConvertGroupApp.PodConvert

