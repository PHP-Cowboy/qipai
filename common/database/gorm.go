package database

import (
	"common/config"
	"common/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// 定义自己的Writer
type SqlWriter struct {
	sqlLog *zap.SugaredLogger
}

// 实现gorm/logger.Writer接口
func (m *SqlWriter) Printf(format string, v ...interface{}) {
	//记录日志
	m.sqlLog.Info(fmt.Sprintf(format, v...))
}

func NewSqlWriter() *SqlWriter {
	l, ok := global.Logger["sql"]

	if !ok {
		panic("sql日志加载失败")
	}

	return &SqlWriter{sqlLog: l}
}

type DbGorm struct {
	*gorm.DB
}

func NewDbGorm() *DbGorm {
	cfg := config.Conf.Database.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	logger := logger2.New(
		NewSqlWriter(),
		logger2.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			Colorful:      true,        //禁用彩色打印
			LogLevel:      logger2.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "d_", // 表名前缀
			SingularTable: true, // 使用单数表名
		},
		Logger: logger,
	})

	if err != nil {
		panic(err)
	}

	return &DbGorm{db}
}
