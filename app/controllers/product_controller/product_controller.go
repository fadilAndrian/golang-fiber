package product_controller

import (
	"github.com/gofiber/fiber/v2"
	"learn-fibergo/app/models"
	"learn-fibergo/database"
	"log"
)

func Create(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		log.Fatal(err.Error())
	}

	query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, name, price`

	if err := database.DB.QueryRow(query, product.Name, product.Price).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error inserting product",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product added",
		"data":    &product,
	})
}

func List(c *fiber.Ctx) error {
	var products []models.Product

	query := `SELECT * FROM products`
	rows, err := database.DB.Query(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error listing products",
			"error":   err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product

		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error listing products",
				"error":   err.Error(),
			})
		}

		products = append(products, product)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": products,
	})
}

func Show(c *fiber.Ctx) error {
	query := `SELECT * FROM products WHERE id = $1`
	id := c.Params("id")

	var product models.Product

	if err := database.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data nof found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": product,
	})
}

func Update(c *fiber.Ctx) error {
	query := `UPDATE products SET name = $1, price = $2 WHERE id = $3`
	id := c.Params("id")

	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err.Error(),
		})
	}

	if err := database.DB.QueryRow(query, product.Name, product.Price, id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated",
		"data":    product,
	})
}
