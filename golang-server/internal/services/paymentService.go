package services

import (
	"golang-server/internal/dto"
	"time"
)

func Payment(paymentRequest dto.PaymentRequest) {
	time.Sleep(3 * time.Second)
}
