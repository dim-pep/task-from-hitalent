package db

import (
	"fmt"
	"log"

	"github.com/dim-pep/task-from-hitalent/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func Conn(conf config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName, conf.DBSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto;")
	log.Println("Успешно подключился к базе данных")
	DbConn = db
}
