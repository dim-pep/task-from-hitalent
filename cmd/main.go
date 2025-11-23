package main

import (
	"log"

	"github.com/dim-pep/task-from-hitalent/config"
	"github.com/dim-pep/task-from-hitalent/internal/db"
)

func main() {
	conf := config.LoadConfig()
	log.Println("Попытка подключения к базе данных")
	db.Conn(conf)
}
