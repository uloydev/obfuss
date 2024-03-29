package libs

import (
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"skripsi.id/obfuss/entities"
)

func NewMysqlConn(logger *zap.Logger) *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: glog.Default.LogMode(glog.Info),
	})
	if err != nil {
		logger.Fatal("error when connect to mysql", zap.Error(err))
	}

	err = db.AutoMigrate(&entities.Todo{})
	if err != nil {
		logger.Fatal("error when migrate table", zap.Error(err))
	}

	logger.Info("mysql database connected")

	return db
}
