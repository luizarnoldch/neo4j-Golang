package dto

type ItemRequest struct {
	ItemID   int64  `json:"id" form:"id"`
	ItemName string `json:"name" form:"name"`
}
