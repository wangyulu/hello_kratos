package data

import (
	stdlog "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hello/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "repo/user")),
	}
	cleanup := func() {
		data.log.Info("closing the data resources")
	}

	return data, cleanup, nil
}

func NewDB(c *conf.Data, sLogger log.Logger) (*gorm.DB, func(), error) {
	log := log.NewHelper(log.With(sLogger, "data", "db"))
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢查询 SQL 阈值
			Colorful:      true,        // 禁用彩色打印
			// IgnoreRecordNotFoundError: false,
			LogLevel: logger.Info, // Log lever
		},
	)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			// SingularTable: true, // 表名是否加
		},
	})

	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		panic("failed to connect database")
	}

	cleanFn := func() {}

	return db, cleanFn, nil
}
