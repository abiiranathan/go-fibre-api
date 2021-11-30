package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abiiranathan/goclinic/database"
	"github.com/abiiranathan/goclinic/models"
	"github.com/abiiranathan/goclinic/validation"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	fmt.Println(db)

	var users []models.User

	db.Find(&users)

	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	db := database.DB

	id := c.Params("id")

	var user models.User

	db.Find(&user, id)

	if user.ID == 0 {
		return c.Status(http.StatusNotFound).JSON("User not found!")
	}

	return c.JSON(user)
}

func GetUserByEmail(c *fiber.Ctx) error {
	db := database.DB
	email := c.Params("email")

	var user models.User

	db.Where("email=?", email).First(&user)

	if user.ID == 0 {
		return c.Status(http.StatusNotFound).JSON("User not found!")
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	errors := validation.ValidateModel(*user)
	if errors != nil {
		return c.JSON(errors)
	}

	db.Preload("Permissions").Create(&user)
	return c.Status(http.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var user models.User

	db.First(&user, id)

	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusBadRequest).Send([]byte(err.Error()))
	}

	// Do not update the id
	userId, _ := strconv.Atoi(id)
	user.ID = uint(userId)

	database.DB.Preload("Permissions").Updates(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var user models.User

	db.First(&user, id)

	db.Delete(&user)

	return c.Status(http.StatusNoContent).JSON(nil)
}
