package external

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ExchangeRateRaw struct {
	Result       int    `json:"RESULT"`          // 1 : 성공, 2 : DATA코드 오류, 3 : 인증코드 오류, 4 : 일일제한횟수 마감
	CurUnit      string `json:"CUR_UNIT"`        // 통화코드
	CurName      string `json:"CUR_NM"`          // 국가/통화명
	Ttb          string `json:"TTB"`             // 전신환(송금) 받으실때
	Tts          string `json:"TTS"`             // 전신환(송금) 보내실때
	DealBasRate  string `json:"DEAL_BAS_R"`      // 매매 기준율
	Bkpr         string `json:"BKPR"`            // 장부가격
	YyEfeeR      string `json:"YY_EFEE_R"`       // 년환가료율
	TenDdEfeeR   string `json:"TEN_DD_EFEE_R"`   // 10일환가료율
	KftcDealBasR string `json:"KFTC_DEAL_BAS_R"` // 서울외국환중개 매매기준율
	KftcBkpr     string `json:"KFTC_BKPR"`       // 서울외국환중개 장부가격
}

func GetExchangeRate(date string) ([]ExchangeRateRaw, error) {

	client := &http.Client{Timeout: 3 * time.Second}

	baseURL := "https://oapi.koreaexim.go.kr/site/program/financial/exchangeJSON"
	openApiKey := "fmg4TI2nZeveC6j7e5bh35v6EtpSCveZ"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	//now := time.Now()
	//targetDate := now
	//if now.Hour() < 11 {
	//	targetDate = now.AddDate(0, 0, -1)
	//}

	q := req.URL.Query()
	q.Add("authkey", openApiKey)
	q.Add("searchdate", date)
	//q.Add("searchdate", "20251119")
	q.Add("data", "AP01")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var parsed []ExchangeRateRaw
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}
