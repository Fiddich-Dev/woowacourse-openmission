package dto

import (
	"golang-server/internal/models"
)

type PaymentRequest struct {
	GoodsList     []models.Goods `json:"goods_list"`
	PaymentAmount string         `json:"payment_amount"`
}

type GoodsListPartitionRequest struct {
	GoodsList []models.Goods `json:"goods_list"`
}
