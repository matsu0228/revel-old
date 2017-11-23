// package db
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "myapp/app/models/entity"
)

func gormConnect() *gorm.DB {
	USER := "dev"
	PASS := "dev1709"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "go_tutorial"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type Book struct {
	Id     int    `json:id`
	Name   string `json:name`
	Author string `json:author`
}

// func main() {
// 	db, err := gorm.Open("mysql", "dev:dev1709@/go_tutorial?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// }

func main() {
	db := gormConnect()
	// var book Book
	var count = 0

	// insert
	// ---------------------------
	new_book := Book{Name: "Sikisai", Author: "Haruki"}
	// if db.NewRecord(new_book) // => returns `true` as primary key is blank
	db.Create(&new_book)

	// select
	// ---------------------------
	bookEx := []Book{}
	// bookEx.Id = 1
	db.Find(&bookEx).Count(&count)
	// db.Last(&bookEx).Count(&count)
	// db.First(&bookEx).Count(&count)
	if count == 0 {
		fmt.Println("該当レコードなし")
	} else {
		for i := 0; i < len(bookEx); i++ {
			fmt.Println("book:" + fmt.Sprint(bookEx[i].Id) + " =id/ " + bookEx[i].Name + " / " + bookEx[i].Author)
		}
	}
	defer db.Close()
}
