package project

import (
	"github.com/celestix/autodeployer_api/api/project/advanced"
	"github.com/celestix/autodeployer_api/api/project/deploy"
	"github.com/celestix/autodeployer_api/api/project/graph"
	"github.com/celestix/autodeployer_api/api/project/info"
	"github.com/celestix/autodeployer_api/api/project/list"
	"github.com/celestix/autodeployer_api/api/project/new"
	"github.com/celestix/autodeployer_api/api/project/resources"
	"github.com/celestix/autodeployer_api/middleware"
	"github.com/gin-gonic/gin"
)

func Load(rg *gin.RouterGroup) {
	rg = rg.Group("/project")
	rg.Use(middleware.Auth)
	list.Load(rg)
	advanced.Load(rg)
	deploy.Load(rg)
	graph.Load(rg)
	info.Load(rg)
	new.Load(rg)
	resources.Load(rg)
}
