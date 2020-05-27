package routers

import (
	"database/sql"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/pprof"
	"github.com/zainokta/bookstore_fiber/controllers"
	"github.com/zainokta/bookstore_fiber/repositories"
)

func Router(db *sql.DB) *fiber.App {
	router := fiber.New()
	router.Use(pprof.New())
	// router.Use(logger.New())
	router.Use(helmet.New())

	bookRepo := repositories.NewBookRepo(db)
	bookHandler := controllers.NewBookHandler(bookRepo)

	router.Get("/", bookHandler.IndexBook)
	router.Get("/:id", bookHandler.ShowBook)
	router.Post("/", bookHandler.StoreBook)
	router.Put("/:id", bookHandler.UpdateBook)
	router.Delete("/:id", bookHandler.DeleteBook)

	return router
}
