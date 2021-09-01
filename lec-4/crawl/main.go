package main

import (
	crawler "github.com/trandtung1209/crawl/crawler"

	db "github.com/trandtung1209/crawl/database"
	"github.com/trandtung1209/crawl/entity"
)

func main() {
	con:= db.ConnectDB()
	con.Migrator().HasTable(&entity.Film{})
	con.Migrator().DropTable(&entity.Film{})
	con.Migrator().CreateTable(&entity.Film{})

	crawler.Crawl()
}