package list

import (
	"github.com/celestix/autodeployer_api/internal/db"
	"github.com/gin-gonic/gin"
)

func listCallback(ctx *gin.Context) {
	projs, err := db.ListProjects()
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, projs)
	return
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/list", listCallback)
}
