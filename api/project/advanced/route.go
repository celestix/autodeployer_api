package advanced

import "github.com/gin-gonic/gin"

func advancedCallback(ctx *gin.Context) {

}

func Load(rg *gin.RouterGroup) {
	rg.GET("/advanced", advancedCallback)
}
