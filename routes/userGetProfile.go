package routes

import (
	"net/http"

	"github.com/gofiber/fiber"
	database "github.com/jin-wk/fiber-mysql/database/user"
)

// RequestID : Request id for get profile
type RequestID struct {
	ID string `json:"ID"`
}

// GetProfile : get user profile
func GetProfile(c *fiber.Ctx) error {
	b := new(RequestID)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	if len(b.ID) < 1 {
		c.SendString(b.ID)
		c.SendString("Parameter ID is required")
		c.SendStatus(http.StatusBadRequest)
		return nil
	}
	profile, err := database.SearchProfile(b.ID)
	if profile.IsEmpty() {
		c.SendString("Not Found User")
		c.SendStatus(http.StatusNotFound)
		return nil
	}
	if err != nil {
		c.SendString("Error Occureed" + err.Error())
		c.SendStatus(http.StatusBadRequest)
		return nil
	}
	if err := c.JSON(profile); err != nil {
		c.Status(500).SendString(err.Error())
		return nil
	}
	c.Accepts("application/json")
	c.SendStatus(http.StatusOK)
	return nil
}
