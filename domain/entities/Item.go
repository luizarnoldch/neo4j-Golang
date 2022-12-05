package entities

import "github.com/luizarnoldch/neo4j-Golang/domain/dto"

type Item struct {
	ItemID   int64  `db:"id"`
	ItemName string `db:"name"`
}

func (i Item) ToItemResponse() *dto.ItemResponse {
	return &dto.ItemResponse{i.ItemID, i.ItemName}
}

func NewItem(ItemID int64, ItemName string) Item {
	return Item{ItemID, ItemName}
}
