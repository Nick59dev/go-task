package main

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
  "github.com/antchfx/htmlquery"
	"fmt"
)

type Product struct {
	gorm.Model
	// id         uint   `gorm:"<-:create"` created automatically in psql
	name       string `gorm:"<-:create"`
	url        string `gorm:"<-:create"`
	url_image  string `gorm:"<-:create"`
	price      string `gorm:"<-:create"`
}

func main() {
	var product Product;

  dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai" // enter your data here
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection trouble. Shutting down.")
	}

  sqlDB, err := sql.Open("pgx", "mydb_dsn")
  gormDB, err := gorm.Open(postgres.New(postgres.Config{
    Conn: sqlDB,
  }), &gorm.Config{})

	// creating the doc file
	doc, err := htmlquery.LoadURL("http://www.ozon.ru/category/moloko-9283")
	if err != nil {
		fmt.Println(doc, product)
		panic("error occured.")
	}

	// parsing necessary <div>
	list := htmlquery.Find(doc, "//*[@class='j1u ju2']")
	for _, n := range list {
		// fmt.Println(htmlquery.InnerText(n)) // htmlquery.SelectAttr(n, "class"))

		// choosing necessary image link in <a>
		elem, err := htmlquery.Find(n, "//a[@class='tj3 tile-hover-target']/@href")
		var img string = htmlquery.SelectAttr(elem, "href")
		product.url = img

		elem, err := htmlquery.Find(n, "//*/img[@class='ui-p4']/@href")
		img = htmlquery.SelectAttr(elem, "href")
		product.url_image = img

		// choosing necessary name for the product
		elem, err := htmlquery.Find(n, "//*[@class='d9m m9d dn0 n1d tsBodyL sj5']/span")
		img = htmlquery.InnerText(elem)
		product.name = img

		elem, err := htmlquery.Find(n, "//*[@class='ui-q5 ui-q9']")
		img = htmlquery.InnerText(elem)
		product.price = img

		result := db.Create(&product)
		if result.Error != nil {
			panic("database inserting error. Shutting down...")
		}
	}

	// fmt.Println(doc)
}
