package database

import (
	"authapp/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func InitDatabase() {
	cfg := config.Get()

	dbResourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if cfg.DBDriver == "sqlite3" {
		dbResourceName = cfg.DBHost
	}

	db, err := gorm.Open(cfg.DBDriver, dbResourceName)
	if err != nil {
		panic(err)
	}

	DB = db

	// set gorm configuration
	DB.LogMode(true)
	DB.SingularTable(false)

	err = DB.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"`id` bigint(20) NOT NULL AUTO_INCREMENT, " +
		"phonenumber varchar(13) not null, " +
		"name varchar(50) not null, " +
		"password varchar(50) not null, " +
		"role varchar(50) not null, " +
		"created_at datetime not null default current_timestamp, " +
		"updated_at datetime not null default current_timestamp, " +
		"deleted_at datetime default null," +
		"PRIMARY KEY (`id`)" +
		")").Error
	if err != nil {
		panic(err)
	}

}
