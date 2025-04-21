package repo

import (
	"github.com/celestix/autodeployer_api/api/repo/branches"
	"github.com/gin-gonic/gin"
)

func Load(rg *gin.RouterGroup) {
	rg = rg.Group("/repo")
	branches.Load(rg)
}
