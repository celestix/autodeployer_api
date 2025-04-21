package new

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/celestix/autodeployer_api/internal/core"
	"github.com/celestix/autodeployer_api/internal/db"
	"github.com/gin-gonic/gin"
)

type NewProjectRequest struct {
	Name                 string                     `json:"name"`
	Branch               string                     `json:"branch"`
	Type                 db.ProjectType             `json:"type"`
	RepoUrl              string                     `json:"repo_url"`
	DeploymentInfo       core.ProjectDeploymentInfo `json:"deployment_info"`
	EnvironmentVariables [][]string                 `json:"environment_variables"`
}

var re = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?github\.com\/([^\/]+)\/([^\/]+)`)

func newCallback(ctx *gin.Context) {
	gho := ctx.GetString("gho")
	var v NewProjectRequest
	err := ctx.BindJSON(&v)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to process json body: %s", err)})
		return
	}
	var envVars = make([]db.ProjectEnvVar, len(v.EnvironmentVariables))
	for i, envVar := range v.EnvironmentVariables {
		if len(envVar) != 2 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "environment variables must be a list of key-value pairs"})
			return
		}
		envVars[i] = db.ProjectEnvVar{Key: envVar[0], Value: envVar[1]}
	}
	buf, err := json.Marshal(v.DeploymentInfo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to marshal deployment info: %s", err)})
		return
	}
	match := re.FindStringSubmatch(v.RepoUrl)
	if match == nil || len(match) != 3 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid repo url"})
		fmt.Println(match)
		fmt.Println(v.RepoUrl)
		return
	}
	repoOwner := match[1] // match[0] contains entire match
	repoName := match[2]
	err = db.AddProject(v.Name, v.Type, repoOwner, repoName, v.Branch, gho, buf, envVars)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to add project: %s", err)})
		return
	}
	err = core.CloneRepository(v.Name, repoOwner, repoName, v.Branch, gho)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to clone repository: %s", err)})
		return
	}
	var args = make([]string, len(v.EnvironmentVariables))
	for i, arg := range v.EnvironmentVariables {
		args[i] = fmt.Sprintf("%s=%s", arg[0], arg[1])
	}
	err = core.DeployProject(v.DeploymentInfo, &core.DeployParams{
		RepoOwner:   repoOwner,
		RepoName:    repoName,
		EnvVars:     args,
		ProjectName: v.Name,
	})
	if err != nil {
		log.Println("Failed to deploy:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to deploy project: %s", err)})
		return
	}
	fmt.Println(v)
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
	return
}

func Load(rg *gin.RouterGroup) {
	rg.POST("/new", newCallback)
}
