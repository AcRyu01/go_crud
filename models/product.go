package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	ImageUrl string `json:"image_url"`
}

func GetProduct(db *gorm.DB, c *fiber.Ctx) error {
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(products)
}

func GetProductByID(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func CreateProduct(db *gorm.DB, c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return err
	}
	result := db.Create(&product)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func UpdateProduct(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(Product)
	result1 := db.First(&product, id)
	if result1.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := c.BodyParser(product); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result2 := db.Save(&product)
	if result2.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}

func DeleteProduct(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	result := db.Delete(&Product{}, id)
	if result.Error != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.SendString("Book successfully deleted")
}
