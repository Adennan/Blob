package global

import (
	"Blob/config"
	"Blob/utils/logger"

	"github.com/jinzhu/gorm"
)

var (
	Server   *config.Server
	App      *config.App
	Database *config.Database

	DB     *gorm.DB
	Logger *logger.Logger
)
