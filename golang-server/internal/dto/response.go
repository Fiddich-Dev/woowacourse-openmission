package dto

import (
	"golang-server/internal/models"
)

type DiscountPolicyByResponse struct {
	Threshold      string `json:"threshold"`
	DiscountAmount string `json:"discount_amount"`
}

type ExchangeRateResponse struct {
	Result  int    `json:"result"`   // 1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감
	CurUnit string `json:"cur_unit"` // "USD"
	Ttb     string `json:"ttb"`      // 팔떄 가격
	Tts     string `json:"tts"`      // 살떄 가격
	CurNm   string `json:"cur_nm"`   // "미국 달러"
}

type GoodsListPartitionResponse struct {
	GoodsList   []models.Goods `json:"goods_list"`
	BeforePrice string         `json:"before_price"`
	AfterPrice  string         `json:"after_price"`
}
