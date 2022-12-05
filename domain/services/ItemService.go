package services

import (
	"errors"
	"github.com/luizarnoldch/neo4j-Golang/domain/dto"
	"github.com/luizarnoldch/neo4j-Golang/domain/repositories"
)

type ItemService interface {
	GetAllItems() ([]dto.ItemResponse, error)
	PostItem(request dto.ItemRequest) (any, error)
}

type DefaultItemService struct {
	db repositories.ItemRepository
}

func NewItemService(db repositories.ItemRepository) DefaultItemService {
	return DefaultItemService{db}
}

func (s DefaultItemService) GetAllItems() ([]dto.ItemResponse, error) {
	results, errDB := s.db.FindAllPersonas()
	if errDB != nil {
		return nil, errDB
	}
	response := make([]dto.ItemResponse, 0)

	for _, result := range results {
		response = append(response, *result.ToItemResponse())
	}

	return response, nil
}

func (s DefaultItemService) PostItem(request dto.ItemRequest) (any, error) {
	_, err := s.db.SavePersonas(request.ItemID, request.ItemName)
	if err != nil {
		return nil, errors.New("can't save item")
	}
	return "guardado correcto", nil
}
