package auth

import (
	"log"

	"github.com/celestix/autodeployer_api/config"
	"github.com/celestix/autodeployer_api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v70/github"
)

func authCallback(ctx *gin.Context) {
	tmpCode := ctx.Query("code")
	if tmpCode == "" {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "code is required"})
		return
	}
	token, err := getAccessTokenGithub(config.Data.GhOauthClientId, config.Data.GhOauthClientSecret, tmpCode)
	if err != nil {
		log.Println("failed to getAccessTokenGithub:", err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": "failed to get access token"})
		return
	}
	gClient := github.NewClient(nil).WithAuthToken(token.AccessToken)
	user, _, err := gClient.Users.Get(ctx, "")
	if err != nil {
		log.Println("failed to get user:", err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": "failed to get user"})
		return
	}
	log.Println("user:", user)
	var name string = "user"
	if user.Name != nil {
		name = *user.Name
	} else if user.Email != nil {
		name = *user.Email
	} else if user.Login != nil {
		name = *user.Login
	}
	jwt, err := utils.GenerateJWT(name, token.AccessToken)
	if err != nil {
		log.Println("failed to generate jwt:", err)
		ctx.AbortWithStatusJSON(500, gin.H{"error": "failed to generate jwt"})
		return
	}
	ctx.String(200, jwt)
}

func Load(rg *gin.RouterGroup) {
	rg.GET("/auth/github", authCallback)
}
