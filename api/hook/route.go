package hook

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/celestix/autodeployer_api/internal/core"
	"github.com/celestix/autodeployer_api/internal/db"
	"github.com/gin-gonic/gin"
)

type HookRequest struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func hookCallback(ctx *gin.Context) {
	var hr HookRequest
	err := ctx.BindJSON(&hr)
	if err != nil {
		log.Println("HOOK:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to process json body"})
		return
	}
	repo := strings.Split(hr.Repository.FullName, "/")
	repoOwner := repo[0]
	repoName := repo[1]
	project, err := db.GetProjectByName(repoOwner, repoName)
	if err != nil {
		log.Println("HOOK-ERR:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get project"})
		return
	}
	err = core.PullRepository(project.Name, project.Gho)
	if err != nil {
		log.Println("HOOK-ERR:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to pull repository"})
		return
	}
	deployment, err := core.GetProjectDeploymentInfo(project.Type, project.Deployment)
	if err != nil {
		log.Println("HOOK-ERR:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unknown failure"})
		return
	}
	err = deployment.Stop(project.Name)
	if err != nil {
		log.Println("HOOK-ERR:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "unknown failure"})
		return
	}
	var args = make([]string, len(project.EnvVars))
	for i, arg := range project.EnvVars {
		args[i] = fmt.Sprintf("%s=%s", arg.Key, arg.Value)
	}
	err = deployment.Deploy(&core.DeployParams{
		ProjectName: project.Name,
		RepoName:    project.RepoName,
		RepoOwner:   project.RepoOwner,
		EnvVars:     args,
	})
	fmt.Println("HOOK-ERR:", err)
	ctx.String(200, "ok")
	return
}

func Load(rg *gin.RouterGroup) {
	rg.POST("/hook", hookCallback)
}
