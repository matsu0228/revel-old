// package dbinit
package main

import (
	"fmt"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
	// "strings"
	// "myapp/app/models/entity"
)

// func gormConnect() *gorm.DB {
func gormConnect() {
	// USER := "dev"
	// PASS := "dev1709"
	// PROTOCOL := "tcp(127.0.0.1:3306)"
	// DBNAME := "go_tutorial"
	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	connectionString := getConnectionString()
	// ---------------------------
	fmt.Println("con:" + connectionString)
	// db, err := gorm.Open("mysql", connectionString)
	//
	// if err != nil {
	// 	panic(err.Error())
	// }
	// return db
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

func getConfigParam(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getConfigParam("db.host", "")
	port := getConfigParam("db.port", "3306")
	user := getConfigParam("db.user", "")
	pass := getConfigParam("db.password", "")
	dbname := getConfigParam("db.name", "")
	protocol := getConfigParam("db.protocol", "tcp")
	// dbargs := getConfigParam("db.args", " ")
	//
	// if strings.Trim(dbargs, " ") != "" {
	// 	dbargs = "?" + dbargs
	// } else {
	// 	dbargs = ""
	// }
	dbargs := ""
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func main() {
	connectionString := getConnectionString()
	fmt.Println("con:" + connectionString)
	// db := gormConnect()
	// // var book Book
	// var count = 0
	//
	// // insert
	// // ---------------------------
	// new_book := Book{Name: "Sikisai", Author: "Haruki"}
	// // if db.NewRecord(new_book) // => returns `true` as primary key is blank
	// db.Create(&new_book)
	//
	// // select
	// // ---------------------------
	// bookEx := []Book{}
	// // bookEx.Id = 1
	// db.Find(&bookEx).Count(&count)
	// // db.Last(&bookEx).Count(&count)
	// // db.First(&bookEx).Count(&count)
	// if count == 0 {
	// 	fmt.Println("該当レコードなし")
	// } else {
	// 	for i := 0; i < len(bookEx); i++ {
	// 		fmt.Println("book:" + fmt.Sprint(bookEx[i].Id) + " =id/ " + bookEx[i].Name + " / " + bookEx[i].Author)
	// 	}
	// }
	// defer db.Close()
}
