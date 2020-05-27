package controllers

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/zainokta/bookstore_fiber/models"
)

func (b *BookHandler) IndexBook(ctx *fiber.Ctx) {
	// books, err := b.bookRepo.All()
	// if err != nil {
	// 	ctx.Status(500).Send(err)
	// 	return
	// }
	// resp := map[string][]models.Book{"data": books}

	// if err := ctx.JSON(resp); err != nil {
	// 	ctx.Status(500).Send(err)
	// 	return
	// }
	ctx.Write("Hello")
}

func (b *BookHandler) ShowBook(ctx *fiber.Ctx) {
	idParam := ctx.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	book, err := b.bookRepo.Find(id)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	resp := map[string]models.Book{"data": book}

	if err := ctx.JSON(resp); err != nil {
		ctx.Status(500).Send(err)
		return
	}
}

func (b *BookHandler) StoreBook(ctx *fiber.Ctx) {
	model := new(models.Book)
	if err := ctx.BodyParser(model); err != nil {
		ctx.Status(400).Send(err)
		return
	}

	err := b.bookRepo.Create(model)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	ctx.JSON(map[string]string{"message": "success"})
}

func (b *BookHandler) UpdateBook(ctx *fiber.Ctx) {
	idParam := ctx.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	model := new(models.Book)
	if err := ctx.BodyParser(model); err != nil {
		ctx.Status(400).Send(err)
		return
	}

	err = b.bookRepo.Update(id, model)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	ctx.JSON(map[string]string{"message": "success"})
}

func (b *BookHandler) DeleteBook(ctx *fiber.Ctx) {
	idParam := ctx.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	err = b.bookRepo.Delete(id)
	if err != nil {
		ctx.Status(500).Send(err)
		return
	}

	ctx.JSON(map[string]string{"message": "success"})
}
