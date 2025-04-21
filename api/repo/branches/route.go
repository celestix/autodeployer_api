package branches

import (
	"regexp"

	"github.com/celestix/autodeployer_api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v70/github"
)

var re = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?github\.com\/([^\/]+)\/([^\/]+)`)

func callback(ctx *gin.Context) {
	gho := ctx.GetString("gho")
	url := ctx.Query("url")
	if url == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "url is required"})
		return
	}
	match := re.FindStringSubmatch(url)
	if match == nil || len(match) != 3 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid url"})
		return
	}
	owner := match[1] // match[0] contains entire match
	repo := match[2]
	client := github.NewClient(nil).WithAuthToken(gho)
	repoInfo, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "failed to get repo info"})
		return
	}
	var branches = make([]string, len(repoInfo))
	for i, branch := range repoInfo {
		branches[i] = branch.GetName()
	}
	ctx.JSON(200, gin.H{"branches": branches})
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/branches", middleware.Auth, callback)
}
