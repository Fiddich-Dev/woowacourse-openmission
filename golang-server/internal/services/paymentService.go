package services

import (
	"golang-server/internal/dto"
	"math/rand"
	"time"
)

func Payment(paymentRequest dto.PaymentRequest) {
	n := rand.Intn(3)
	delay := time.Duration(n+1) * time.Second
	time.Sleep(delay)
	//time.Sleep(3 * time.Second)
}
