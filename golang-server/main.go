package main

import (
	"encoding/json"
	"golang-server/model"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"content": "ok",
		})
	})

	r.POST("/partition", PartitionService)

	r.GET("/exchange-rate", GetExchangeRate)

	r.GET("/coupon-discount-policy", model.CouponDiscountPolicyService)
	r.GET("/card-discount-policy", model.CardDiscountPolicyService)

	r.Run()
}

func GetExchangeRate(c *gin.Context) {
	client := &http.Client{Timeout: 3 * time.Second}

	baseURL := "https://oapi.koreaexim.go.kr/site/program/financial/exchangeJSON"
	openApiKey := "fmg4TI2nZeveC6j7e5bh35v6EtpSCveZ"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	q := req.URL.Query()
	q.Add("authkey", openApiKey)
	q.Add("data", "AP01")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var parsed []ExchangeRate
	json.Unmarshal(body, &parsed)

	var response ExchangeRate
	for idx, exchageRate := range parsed {
		if exchageRate.CurUnit == "USD" {
			response = parsed[idx]
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

type ExchangeRate struct {
	Result  int    `json:"result"`   // 1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감
	CurUnit string `json:"cur_unit"` // "USD"
	Ttb     string `json:"ttb"`      // 팔떄 가격
	Tts     string `json:"tts"`      // 살떄 가격
	CurNm   string `json:"cur_nm"`   // "미국 달러"
}
