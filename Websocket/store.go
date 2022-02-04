package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	// autoMigrate bool
)

func InitDB() *gorm.DB {

	d, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/room?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully connect to database")
	}
	//defer db.Close()

	// Run Migrations
	d.LogMode(true)
	errors := d.AutoMigrate(&Room{}).Error
	if errors != nil {
		log.Println(errors.Error())
	}
	db = d
	return db
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		err := fmt.Errorf("db is not initialized")
		return nil, err
	}
	return db, nil
}
