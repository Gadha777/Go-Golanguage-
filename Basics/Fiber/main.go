package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Define a struct to hold user input
type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Global slice to store users
var users []UserInput

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/data", func(c *fiber.Ctx) error {
		var user UserInput

		if err := c.BodyParser(&user); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
		}
		users = append(users, user)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User registered successfully",
			"user":    user,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		fmt.Println("Fetching users:", users)
		return c.JSON(users)
	})
	app.Patch("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		var updateUser UserInput
		if err := c.BodyParser(&updateUser); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
		}
		// Loop through users and update the matching user by name
		for i, user := range users {
			if user.Name == name {
				users[i].Email = updateUser.Email
				users[i].Age = updateUser.Age

				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "User updated successfully",
					"user":    users[i],
				})
			}
		}

		// If the user with the given name is not found
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})

	app.Delete("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		// Loop through the users slice to find and remove the user by name
		for i, user := range users {
			if user.Name == name {
				// Remove user from the slice
				users = append(users[:i], users[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message": "User deleted successfully",
				})
			}
		}
		// If user is not found
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	})
	app.Listen(":8080")
}
