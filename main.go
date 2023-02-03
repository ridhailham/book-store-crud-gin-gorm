package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:cdaaptnia404@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	// bookFileRepository := book.NewFileRepository()

	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// bookRepository := book.NewRepository(db)

	//CRUD ALTERNATIF

	// book := book.Book{}
	// book.Title = "Atomic Habit"
	// book.Price = 120000
	// book.Discount = 15
	// book.Rating = 7
	// book.Description = "Book is teach you how to make good habit and break bad habit"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }
	// var book book.Book

	// err = db.Debug().Where("title = ?", "Atomic Habit (Revised Edition").First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error deleting book record")
	// 	fmt.Println("==========================")
	// }

	// ==============================================
	// UPDATE DATA
	// ==============================================
	// book.Title = "Value Investing (Revised Edition"

	// err = db.Save(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8181")

	//urutan flow

	//main
	//handler
	//service
	//repository
	//db
	//mysql
}
