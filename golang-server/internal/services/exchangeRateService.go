package services

import (
	"golang-server/external"
	"golang-server/internal/config"
	"time"
)

func GetExchangeRate() ([]external.ExchangeRateRaw, error) {
	targetDate := time.Now()

	for {
		rates, err := external.GetExchangeRate(targetDate.Format(config.DateFormat))
		if err != nil {
			return nil, err
		}
		if len(rates) > 0 {
			return rates, nil
		}
		targetDate = targetDate.AddDate(0, 0, -1)
	}
}
