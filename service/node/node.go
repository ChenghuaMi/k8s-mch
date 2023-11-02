/**
 * @author mch
 */

package node

import(
	"context"
	"k8s-mch/global"
	node_res "k8s-mch/model/node/response"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type NodeService struct {

}

func(*NodeService) GetNodeDetail(nodeName string) (*node_res.Node,error) {
	nodeK8s, err := global.KubeConfigSet.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil,err
	}
	nodeRes := nodeConvert.GetNodeDetail(*nodeK8s)
	return &nodeRes,nil
}
func(*NodeService) GetNodeList(keyword string) ([]node_res.Node,error) {
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil,err
	}
	nodeList := make([]node_res.Node,0)

	for _,item := range list.Items {
		if strings.Contains(item.Name,keyword) {
			nodeList = append(nodeList,nodeConvert.GetNodeRes(item))
		}


	}
	return nodeList,nil
}