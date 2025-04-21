package db

import (
	"strings"

	"github.com/celestix/autodeployer_api/config"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SESSION *gorm.DB

func Load() error {
	var dialector gorm.Dialector
	if strings.HasPrefix(config.Data.Db_Uri, "sqlite://") {
		dbUri := strings.TrimPrefix(config.Data.Db_Uri, "sqlite://")
		dialector = sqlite.Open(dbUri)
	} else {
		dialector = postgres.Open(config.Data.Db_Uri)
	}
	var err error
	SESSION, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		TranslateError:         true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	dB, err := SESSION.DB()
	if err != nil {
		return err
	}
	dB.SetMaxOpenConns(150)
	SESSION.AutoMigrate(&User{}, &Project{}, &ProjectEnvVar{})
	return nil
}
