package info

import (
	"strconv"

	"github.com/celestix/autodeployer_api/internal/db"
	"github.com/gin-gonic/gin"
)

func infoCallback(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "id must be an integer"})
		return
	}
	proj, err := db.GetProject(uint(id))
	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "project not found"})
		return
	}
	ctx.JSON(200, proj)
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/info", infoCallback)
}
