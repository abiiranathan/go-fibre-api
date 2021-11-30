package routes

import (
	"net/http"
	"strconv"

	"github.com/abiiranathan/goclinic/database"
	"github.com/abiiranathan/goclinic/models"
	"github.com/abiiranathan/goclinic/validation"
	"github.com/gofiber/fiber/v2"
)

func GetPermissions(c *fiber.Ctx) error {
	db := database.DB

	var perms []models.Permission

	db.Find(&perms)

	return c.JSON(perms)
}

func GetPermission(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var perm models.Permission
	db.Find(&perm, id)

	if perm.ID == 0 {
		return c.Status(http.StatusNotFound).JSON("Permission not found!")
	}

	return c.JSON(perm)
}

func CreatePermission(c *fiber.Ctx) error {
	db := database.DB
	perm := new(models.Permission)

	if err := c.BodyParser(perm); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	errors := validation.ValidateModel(*perm)

	if errors != nil {
		return c.JSON(errors)
	}

	db.Create(&perm)
	return c.Status(http.StatusCreated).JSON(perm)
}

func UpdatePermission(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var perm models.Permission

	db.First(&perm, id)

	if err := c.BodyParser(perm); err != nil {
		c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	// Do not update the id
	permId, _ := strconv.Atoi(id)
	perm.ID = int(permId)

	database.DB.Updates(&perm)
	return c.JSON(perm)
}

func DeletePermission(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var perm models.Permission

	db.First(&perm, id)
	db.Delete(&perm)

	return c.Status(http.StatusNoContent).JSON(nil)
}
