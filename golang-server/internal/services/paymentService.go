package services

import (
	"golang-server/internal/dto"
	"time"
)

func Payment(paymentRequest dto.PaymentRequest) {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//sleepDuration := time.Duration(r.Intn(3-1+1)+1) * time.Second
	//time.Sleep(sleepDuration)
	time.Sleep(3 * time.Second)
}
