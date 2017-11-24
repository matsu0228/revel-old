package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
	"strings"
	// "myapp/app/models/db" //not work
)

type App struct {
	*revel.Controller
}

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
	dbargs := getConfigParam("db.args", " ")
	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func gormConnect() *gorm.DB {
	connectionString := getConnectionString()
	db, err := gorm.Open("mysql", connectionString)

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

func getBooks() string {
	db := gormConnect()
	// var book Book
	var count = 0
	var output = ""

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
			output += "\nbook:" + fmt.Sprint(bookEx[i].Id) + " =id/ " + bookEx[i].Name + " / " + bookEx[i].Author
		}
	}
	defer db.Close()
	return output
}

// ----------------------------------------------------------------------------
func (c App) Index() revel.Result {
	// return c.Render()
	// connectionString := getConnectionString()
	book_data := getBooks()
	greeting := "Aloha World/" + book_data
	return c.Render(greeting)
}
func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
