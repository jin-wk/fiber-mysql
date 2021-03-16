package handlers

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/jin-wk/fiber-mysql/database"
	"github.com/jin-wk/fiber-mysql/models"
)

// GetUser godoc
// @Summary Get a User Profile
// @Description Get User Profile by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.ResponseModel{data=models.User}
// @Failure 404 {object} models.ResponseModel{}
// @Failure 503 {object} models.ResponseModel{}
// @Router /api/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)

	if err := database.DB.First(&user, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(models.Response(2000, "Not Found", nil))
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(models.Response(3000, err.Error(), nil))
		}
	}
	return c.JSON(models.Response(1000, "Success", *user))
}
