package main

import (
	"context"
	"fmt"
	"log"

	"github.com/celestix/autodeployer_api/api"
	"github.com/celestix/autodeployer_api/config"
	"github.com/celestix/autodeployer_api/internal/db"
	"github.com/celestix/autodeployer_api/pkg/utils"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalln("Failed to load config: ", err)
	}
	err = db.Load()
	if err != nil {
		log.Fatalln("Failed to load db: ", err)
	}
	startFunc := api.Start
	ctx := context.Background()
	utils.Startup(ctx, startFunc, fmt.Sprintf("API Server started on port :%d", config.Data.Port))
}
