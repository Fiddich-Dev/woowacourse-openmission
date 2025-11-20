package services

import (
	"time"
)

type ExchangeRate struct {
	Result  int    `json:"result"`   // 1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감
	CurUnit string `json:"cur_unit"` // "USD"
	Ttb     string `json:"ttb"`      // 팔떄 가격
	Tts     string `json:"tts"`      // 살떄 가격
	CurNm   string `json:"cur_nm"`   // "미국 달러"
}

func GetExchangeRateDate() string {
	now := time.Now()
	targetDate := now
	if now.Hour() < 11 {
		targetDate = now.AddDate(0, 0, -1)
	}
	return targetDate.Format("20060102")
}
