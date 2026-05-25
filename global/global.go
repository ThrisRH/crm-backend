package global

import (
	"go-sea-crm/internal/setting"
	"go-sea-crm/pkg/logger"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Pdb    *gorm.DB
)
