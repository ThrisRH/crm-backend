package init

import (
	"fmt"
	"go-sea-crm/global"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration postgre", global.Config.Postgre.Dbname)
	InitLogger()
	global.Logger.Info("Config Log Ok", zap.String("OK", "success"))
	InitPostgre()
	InitRedis()

	r := InitRouter()

	r.Run()
}
