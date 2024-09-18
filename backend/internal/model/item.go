package model

import "github.com/akekapong78/workflow/internal/constant"

type Item struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	Title    string              `json:"title"`
	Price    float64             `json:"price"`
	Quantity int                 `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
	OwnerID  uint                `json:"owner_id"`
}

type RequestItem struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type RequestUpdateItemStatus struct {
	Status constant.ItemStatus `json:"status"`
}

type ResponseItem struct {
	ID       uint                `json:"id"`
	Title    string              `json:"title"`
	Price    float64             `json:"price"`
	Quantity int                 `json:"quantity"`
	Status   constant.ItemStatus `json:"status"`
}
