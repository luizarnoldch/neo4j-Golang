package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizarnoldch/neo4j-Golang/domain/dto"
	"github.com/luizarnoldch/neo4j-Golang/domain/services"
)

type ItemHandler struct {
	service services.ItemService
}

func (h ItemHandler) GetAllItems(c *fiber.Ctx) error {
	res, err := h.service.GetAllItems()
	//fmt.Println(res)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (h ItemHandler) SaveItem(c *fiber.Ctx) error {
	req := new(dto.ItemRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	res, errService := h.service.PostItem(*req)
	if errService != nil {
		return errService
	}
	return c.JSON(res)
}
