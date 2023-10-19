package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	// "github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

// Will be explained in the future
var (
	addr = flag.String("addr,':'", os.Getenv("PORT"), "") // Port for the project from the .env file
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")

	// flag is used for command-line utilities to accept flags to customize the command’s execution eg in https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
)

func Run() error {
	flag.Parse() // if any of the above key is passed as flag in cmd, then the following  defines how it should be used

	if *addr == ":" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")        // Accessing all html files form views folder to render
	app := fiber.New(fiber.Config{Views: engine}) // Setting app with the html for routing
	app.Use(logger.New())                         // to log all inputs
	app.Use(cors.new())                           // CORS middleware for Fiber that can be used to enable Cross-Origin Resource Sharing with various options

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomviewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/chat")
	app.Get("/stream/:ssuid/chat/websocket")
	app.Get("/stream/:ssuid/viewer/websocket")

	return nil
}
