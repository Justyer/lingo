package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBInfo struct {
	User   string
	Pass   string
	Host   string
	Port   int
	DbName string
}

func Conn(dbi DBInfo) *gorm.DB {
	dbLinkTemp := "%s:%s@tcp(%s:%d)/%s"
	dbLink := fmt.Sprintf(dbLinkTemp,
		dbi.User, dbi.Pass, dbi.Host,
		dbi.Port, dbi.DbName)
	db, err := gorm.Open("mysql", dbLink)
	if err != nil {
		log.Fatal("[connect mysql err]: ", err)
	}
	return db
}
