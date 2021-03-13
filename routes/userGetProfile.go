package routes

import (
	"net/http"

	"github.com/gofiber/fiber"
	database "github.com/jin-wk/fiber-mysql/database/user"
	"github.com/jin-wk/fiber-mysql/models"
)

// RequestID : Request id for get profile
type RequestID struct {
	ID string `json:"ID"`
}

// GetProfile godoc
// @Summary Get a Profile
// @Description Get Profile by ID
// @ID get-profile-by-id
// @tags Profile
// @Accept  json
// @Produce  json
// @Param ID body RequestID true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /profile [post]
func GetProfile(c *fiber.Ctx) error {
	b := new(RequestID)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	if len(b.ID) < 1 {
		c.SendStatus(http.StatusBadRequest)
		return c.JSON(models.Response(400, "Parameter ID is required"))
	}
	profile, err := database.SearchProfile(b.ID)
	if profile.IsEmpty() {
		c.SendStatus(http.StatusNotFound)
		return c.JSON(models.Response(404, "Not Found User"))
	}
	if err != nil {
		c.SendStatus(http.StatusBadRequest)
		return c.JSON(models.Response(400, "Error Occureed"+err.Error()))
	}
	if err := c.JSON(profile); err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(models.Response(500, "Error Occureed"+err.Error()))
	}
	c.Accepts("application/json")
	c.SendStatus(http.StatusOK)
	return nil
}
