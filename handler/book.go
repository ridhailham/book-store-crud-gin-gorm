package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return

	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		return

	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
