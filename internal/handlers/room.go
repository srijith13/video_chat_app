package handlers

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String())) // redirects to room page in html when welcome function is called with uuidas params
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}
	// uuid, suuid, _ := createOrGetRoom(uuid) // function need to be created

	return nil
}

func RoomWebsocket(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return nil
	}
	// _, _, room := createOrGetRoom(uuid) // function need to be created

	return nil
}
