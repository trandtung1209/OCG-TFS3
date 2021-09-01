package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
  dsn := "root:@tcp(127.0.0.1:3306)/crawl"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
    log.Fatal(err)
  }
	return db
}