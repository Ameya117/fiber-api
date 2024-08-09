package routes

import (
	"errors"

	"github.com/Ameya117/fiber-api/database"
	"github.com/Ameya117/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	repoonseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(repoonseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product

	database.Database.Db.Find(&products)

	var responseProducts []ProductSerializer
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)

}

func FindProduct(id int, product *models.Product) error { // helper function

	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("product not found")
	}
	return nil

}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}


func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updatedProduct UpdateProduct
	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	product.Name = updatedProduct.Name
	product.SerialNumber = updatedProduct.SerialNumber


	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct (c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err:= database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(204).SendString("Product deleted successfully")
}