package services

import (
	"golang-server/external"
	"time"
)

func GetExchangeRateDate() string {
	now := time.Now()
	targetDate := now
	if now.Hour() < 11 {
		targetDate = now.AddDate(0, 0, -1)
	}
	return targetDate.Format("20060102")
}

func GetExchangeRate() ([]external.ExchangeRateRaw, error) {
	now := time.Now()
	targetDate := now

	rates, error := external.GetExchangeRate(targetDate.Format("20060102"))
	if error != nil {
		return nil, error
	}

	for len(rates) == 0 {
		targetDate = targetDate.AddDate(0, 0, -1)

		rates, error = external.GetExchangeRate(targetDate.Format("20060102"))
		if error != nil {
			return nil, error
		}
	}

	return rates, nil
}
