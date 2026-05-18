package init

import (
	"go-sea-crm/global"
	"go-sea-crm/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
