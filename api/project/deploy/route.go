package deploy

import "github.com/gin-gonic/gin"

func deployCallback(ctx *gin.Context) {

}

func Load(rg *gin.RouterGroup) {
	rg.POST("/deploy", deployCallback)
}
