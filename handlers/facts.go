package handlers

import (
	"github.com/Junkes887/go-db/database"
	"github.com/Junkes887/go-db/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	fact := []models.Fact{}

	database.DB.Db.Find(&fact)

	return c.Status(200).JSON(fact)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
