package routes

import (
	"time"
	"errors"
	"github.com/Ameya117/fiber-api/database"
	"github.com/Ameya117/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"created_at"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSerializer) Order {
	return Order{
		ID:        orderModel.ID,
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(uint(order.UserRefer), &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	reponseUser := CreateResponseUser(user)
	reponseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, reponseUser, reponseProduct)
	return c.Status(200).JSON(responseOrder)
}


func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	database.Database.Db.Find(&orders)

	var responseOrders []Order
	for _, order := range orders {
		var user models.User
		FindUser(uint(order.UserRefer), &user)
		var product models.Product
		FindProduct(order.ProductRefer, &product)

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)

	if order.ID == 0 {
		return errors.New("order not found")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var user models.User
	FindUser(uint(order.UserRefer), &user)
	var product models.Product
	FindProduct(order.ProductRefer, &product)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func UpdateOrder(c *fiber.Ctx) error {
	var order models.Order
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Save(&order)

	var user models.User
	FindUser(uint(order.UserRefer), &user)
	var product models.Product
	FindProduct(order.ProductRefer, &product)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}

func DeleteOrder(c *fiber.Ctx) error {
	var order models.Order
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err)
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	database.Database.Db.Delete(&order)

	return c.Status(200).JSON("Order deleted")
}