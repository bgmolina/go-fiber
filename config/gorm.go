package config

import (
	"github.com/bgmolina/go-fiber/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	logger := utils.LoggerFn("gorm")
	env := EnvFn()
	var err error

	DB, err = gorm.Open(postgres.Open(env.PGSQL_URI), &gorm.Config{})
	if err != nil {
		logger.Error("[PostgreSQL]: " + err.Error())
	} else {
		logger.Info("[PostgreSQL] Database connected")
	}
}
