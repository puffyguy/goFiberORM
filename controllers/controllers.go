package controllers

import (
	"errors"
	"net/http"

	"github.com/puffyguy/goFiberORM/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

var validate = validator.New()

/*
# Welcome - Greets the user
*/
func Welcome(c *fiber.Ctx) {
	c.Send("Welcome to go fiber")
	models.DB.AutoMigrate(models.Book{})
}

/*
# GetBooks - List all books details from database
*/
func GetBooks(c *fiber.Ctx) {
	var books []models.Book
	findAllRes := models.DB.Find(&books)
	if findAllRes.RowsAffected == 0 {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "no books available"})
		return
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"Data": &books})
}

/*
# GetBook - List deatils of single book from database using provided isbn

Params data
- isbn string
*/
func GetBook(c *fiber.Ctx) {
	isbn := c.Params("isbn")
	var book models.Book
	findOneRes := models.DB.First(&book, "isbn = ?", isbn)
	if findOneRes.Error != nil {
		if errors.Is(findOneRes.Error, gorm.ErrRecordNotFound) {
			c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "No record found for given isbn"})
		}
		return
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"data": &book})
}

/*
# NewBook - Add new book details to the database

JSON data
- name string
- isbn string
- author string
- price int
*/
func NewBook(c *fiber.Ctx) {
	c.Send("Add New Book")
	var book models.Book

	bindErr := c.BodyParser(&book)
	if bindErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"bind error": bindErr.Error()})
		return
	}
	if validateErr := validate.Struct(&book); validateErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"validate error": validateErr.Error()})
		return
	}
	newBook := models.Book{
		Name:   book.Name,
		ISBN:   book.ISBN,
		Author: book.Author,
		Price:  book.Price,
	}
	models.DB.AutoMigrate(&book)
	insertRes := models.DB.Create(&newBook)
	if insertRes.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": insertRes.Error})
		return
	}
	c.Status(http.StatusCreated).JSON(&fiber.Map{"result": newBook})
}

/*
# UpdateBook - Update single book details in the database using provided isbn

Params data
- isbn string

JSON data
- name string
- author string
- price int
*/
func UpdateBook(c *fiber.Ctx) {
	isbn := c.Params("isbn")
	var book models.Book
	bindErr := c.BodyParser(&book)
	if bindErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"bind error": bindErr.Error()})
		return
	}
	if validateErr := validate.Struct(&book); validateErr != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"validate error": validateErr.Error()})
		return
	}
	updateRes := models.DB.Model(&book).Where("isbn = ?", isbn).Updates(&book)
	if updateRes.RowsAffected == 0 {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "cannot update row with given isbn"})
		return
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"data": "book updated successfully"})
}

/*
# DeleteBook - Removes a single book from the database using provided isbn

Params data
- isbn
*/
func DeleteBook(c *fiber.Ctx) {
	isbn := c.Params("isbn")
	var book models.Book
	deleteRes := models.DB.Delete(&book, isbn)
	if deleteRes.RowsAffected == 0 {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"error": "no record found with given isbn to delete"})
		return
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"data": "book deleted successfully"})
}
