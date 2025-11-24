package main

import (
	"log"

	"github.com/dim-pep/task-from-hitalent/config"
	"github.com/dim-pep/task-from-hitalent/internal/db"
	"github.com/dim-pep/task-from-hitalent/internal/web"
)

func main() {
	conf := config.LoadConfig()
	log.Println("Попытка подключения к базе данных")
	db.Conn(conf)
	web.StartWeb(conf)
}
