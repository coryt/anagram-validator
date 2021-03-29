package ports

import (
	"log"

	"github.com/coryt/anagram/internal/application"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func ServeWeb(application application.Application) *fiber.App {
	httpServer := fiber.New()
	go func(httpServer *fiber.App) {
		httpServer.Static("/static", "public/static")
		// httpServer.Use(csrf.New())
		httpServer.Use(requestid.New())
		httpServer.Use(logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n​",
		}))
		httpServer.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:3000, http://localhost:4000",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
		httpServer.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.SendFile("./public/index.html")
		})
		apiRouter := httpServer.Group("api")
		apiRouter.Get("/anagrams/top", GetTopAnagrams(application))
		apiRouter.Post("/anagrams/check", ValidateAnagram(application))

		log.Fatal(httpServer.Listen(":4000"))
	}(httpServer)

	return httpServer
}
