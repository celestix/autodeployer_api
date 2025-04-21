package graph

import "github.com/gin-gonic/gin"

func graphCallback(ctx *gin.Context) {

}

func Load(rg *gin.RouterGroup) {
	rg.GET("/graph", graphCallback)
}
