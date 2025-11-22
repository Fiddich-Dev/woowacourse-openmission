package services

import (
	"golang-server/external"
	"time"
)

func GetExchangeRate() ([]external.ExchangeRateRaw, error) {
	targetDate := time.Now()

	for {
		rates, err := external.GetExchangeRate(targetDate.Format("20060102"))
		if err != nil {
			return nil, err
		}
		if len(rates) > 0 {
			return rates, nil
		}
		targetDate = targetDate.AddDate(0, 0, -1)
	}
}
