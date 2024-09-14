package database

import (
	"common/config"
	"common/logs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	*sqlx.DB
}

// 创建数据库连接
func NewDb() *Db {
	cfg := config.Conf.Database.MysqlConfig

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		logs.Fatal("NewDb failed,err:%s", err.Error())

		return nil
	}
	if err = db.Ping(); err != nil {
		logs.Fatal("db.Ping() failed,err:%s", err.Error())
		return nil
	}
	return &Db{db}
}

// 根据查询语句获取多行数据
func (*Db) GetRows(db *sqlx.DB, query string, result interface{}) error {
	err := db.Select(result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
			return nil
		}
		return err
	}
	return nil
}

// 根据查询语句获取单行数据
func (*Db) GetRow(db *sqlx.DB, query string, result interface{}) error {
	err := db.Get(result, query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found")
			return nil
		}
		return err
	}
	return nil
}

// 执行增删改等操作
func (*Db) Execute(db *sqlx.DB, query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
