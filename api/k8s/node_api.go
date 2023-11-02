/**
 * @author mch
 */

package k8s

import (
	"github.com/gin-gonic/gin"
	"k8s-mch/response"
)

type NodeApi struct {

}

func(*NodeApi) GetNodeDetailOrList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	nodeName := ctx.Query("nodeName")
	if nodeName != "" {
		nodeDetail,err := nodeService.GetNodeDetail(nodeName)
		if err != nil {
			response.FailWithMessage(ctx,err.Error())
			return
		}
		response.SuccessWithDetailed(ctx,"ok",nodeDetail)
	}else{
		nodeList,err := nodeService.GetNodeList(keyword)
		if err != nil {
			response.FailWithMessage(ctx,err.Error())
			return
		}
		response.SuccessWithDetailed(ctx,"ok",nodeList)
	}

}