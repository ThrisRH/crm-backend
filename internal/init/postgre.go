package init

import (
	"fmt"
	"go-sea-crm/global"
	"go-sea-crm/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CheckErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgre() {
	m := global.Config.Postgre

	dsn := "host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	var s = fmt.Sprintf(dsn, m.Host, m.Username, m.Password, m.Dbname, m.Port)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		CheckErrorPanic(err, "InitPostgreSQL initialization error!")
	}
	global.Logger.Info("Initialization PostgreSQL Successfully")
	global.Pdb = db
}

func SetPool() {
	m := global.Config.Postgre

	postgresDb, err := global.Pdb.DB()
	if err != nil {
		fmt.Printf("Postgree error: %s::", err)
	}

	postgresDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	postgresDb.SetMaxOpenConns(m.MaxOpenConns)
	postgresDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func miggrateTables() {
	err := global.Pdb.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		global.Logger.Error("AutoMigrate error", zap.Error(err))
	}
}
