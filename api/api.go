package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/celestix/autodeployer_api/api/auth"
	"github.com/celestix/autodeployer_api/api/dashboard"
	"github.com/celestix/autodeployer_api/api/hook"
	"github.com/celestix/autodeployer_api/api/project"
	"github.com/celestix/autodeployer_api/api/repo"
	"github.com/celestix/autodeployer_api/config"
	"github.com/gin-gonic/gin"
)

func loadAPIRoutes(r *gin.Engine) {
	apiRouterGroup := r.Group("/api")
	dashboard.Load(apiRouterGroup)
	hook.Load(apiRouterGroup)
	project.Load(apiRouterGroup)
	auth.Load(apiRouterGroup)
	repo.Load(apiRouterGroup)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight request, return immediately
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func Start() error {
	if config.Data.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	loadAPIRoutes(r)

	log.Println("[ENGINE] Main - Started")
	return r.Run(fmt.Sprintf(":%d", config.Data.Port))
}
