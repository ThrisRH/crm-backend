package init

import (
	"fmt"
	"go-sea-crm/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration postgre", global.Config.Postgre.Dbname)
	InitLogger()
	InitPostgre()
	InitRedis()

	r := InitRouter()

	r.Run()
}
