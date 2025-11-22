package handlers

import (
	"golang-server/external"
	"golang-server/internal/dto"
	"golang-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExchangeRate struct {
	Result  int    `json:"result"`   // 1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감
	CurUnit string `json:"cur_unit"` // "USD"
	Ttb     string `json:"ttb"`      // 팔떄 가격
	Tts     string `json:"tts"`      // 살떄 가격
	CurNm   string `json:"cur_nm"`   // "미국 달러"
}

func GetUsdRateHandler(c *gin.Context) {
	date := services.GetExchangeRateDate()

	rates, err := external.GetExchangeRate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail(http.StatusBadGateway, "환율을 가져오지 못했습니다."))
		return
	}

	// USD만 추출
	var response ExchangeRate
	for _, r := range rates {
		if r.CurUnit == "USD" {
			response = ExchangeRate{
				r.Result,
				r.CurUnit,
				r.Ttb,
				r.Tts,
				r.CurName,
			}
			break
		}
	}

	c.JSON(http.StatusOK, dto.Success(response))
}
