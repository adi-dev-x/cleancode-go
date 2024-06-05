package db

import (
	"fmt"
	"log"
	"myproject/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPGDB(cnf config.Config) *gorm.DB {
	fmt.Println(cnf)

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", cnf.PGHost, cnf.PGUserName, cnf.PGPassword, cnf.PGDBName, cnf.PgPort)
	fmt.Println("this is the database ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err, err.Error(), "driver name", cnf.PgDriverName, "postgres url", dsn)
	}

	if err != nil {
		log.Fatal("not connected to postgres db: ", err.Error())
	}

	log.Println("connected to postgres db successfully!")
	DB = db
	return db
}
