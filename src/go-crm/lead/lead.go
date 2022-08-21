package lead

import (
	"github.com/gofiber/fiber/v2"
	"go-crm/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func init() {
	database.Connect()
	db := database.DB
	db.AutoMigrate(&Lead{})
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DB
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DB
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.SendStatus(503)
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with id")
	}
	db.Delete(&lead)
	return c.SendString("Successfully deleted!")
}
