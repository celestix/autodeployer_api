package dashboard

import (
	"github.com/celestix/autodeployer_api/api/dashboard/graph"
	"github.com/gin-gonic/gin"
)

func Load(rg *gin.RouterGroup) {
	rg = rg.Group("/dashboard")
	graph.Load(rg)
}
