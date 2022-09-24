package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:Smith@2022@tcp(127.0.0.1:3306)/db_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func main() {
	fmt.Println("Testing GORM")

	fmt.Println("Successfully connect to the MySQL Database")

	// fmt.Println("Server is listening :8080")
	// ProductManagement()
	BlogManagement()

}

// Product table structure
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ProductManagement() {
	//Product table migration
	db.AutoMigrate(&Product{})

	//Create
	// db.Create(&Product{Code: "A-111", Price: 200})

	//Read with ID
	var product Product
	db.First(&product, 3) // find product with id
	fmt.Println(product)

	//Read with Column Code Name
	db.First(&product, "code = ?", "A-112")
	fmt.Println(product)

	//Read with Column Price Name
	db.First(&product, "price = ?", "200")
	fmt.Println(product)

	//Update the product
	db.Model(&product).Update("price", 250)

	//Update the multiple value
	// db.Model(&product).Updates(Product{Price: 200, Code: "A-116"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "A-117"})

	//Delete - delete the product
	db.Delete(&product, 1)

}

//Embedded Struct

type Author struct {
	gorm.Model
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

func BlogManagement() {
	db.AutoMigrate(&Blog{})

	//create
	// author := Author{
	// 	Name:  "Skoo",
	// 	Email: "sko@gmial.vd",
	// }
	// blog := Blog{Upvotes: 3, Author: author}
	// db.Create(&blog)

	//db.Omit("Email").Create(&blog) //skip email

	//Batch Insert
	// author1 := Author{
	// 	Name:  "author1",
	// 	Email: "author1@gmial.vd",
	// }
	// author2 := Author{
	// 	Name:  "author2",
	// 	Email: "author2@gmial.vd",
	// }
	// author3 := Author{
	// 	Name:  "author3",
	// 	Email: "author3@gmial.vd",
	// }
	// manyauthors := []Blog{{Author: author1}, {Author: author2}, {Author: author3}}
	// db.Create(&manyauthors)
	fmt.Println("----------Find Method----------")
	var blog Blog
	db.Find(&blog)
	fmt.Println(blog)
	fmt.Println("----------Take Method----------")
	var tblog Blog
	db.Take(&tblog)
	fmt.Println(tblog)
	fmt.Println("----------Last Method----------")
	var lblog Blog
	result := db.Last(&lblog)
	fmt.Println(lblog)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	fmt.Println("----------Use Model ----------")
	mresult := map[string]interface{}{}
	db.Model(&Blog{}).First(&mresult)
	fmt.Println(mresult)
	fmt.Println("----------Find with many id ----------")
	var blogs Blog
	db.Find(&blogs)
	fmt.Println(blogs)
	fmt.Println("----------Where Conditions ----------")
	var filtname Blog
	db.Where("name = ?", "author3").Find(&filtname)
	fmt.Println(filtname)
	//or
	fmt.Println("----------Where Conditions Inline----------")
	var filtinname Blog
	db.Find(&filtinname, "name = ?", "author2")
	fmt.Println(filtinname)
}
