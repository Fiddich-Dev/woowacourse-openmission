package payment

import (
	"golang-server/model"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	Goods         []model.Goods `json:"goods"`
	PaymentAmount string        `json:"payment_amount"`
}

func PaymentService(c *gin.Context) {
	var request PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleepDuration := time.Duration(r.Intn(3-1+1)+1) * time.Second
	time.Sleep(sleepDuration)

	c.JSON(http.StatusOK, gin.H{
		"content": request.Goods,
	})
}

func PaymentAllService(c *gin.Context) {
	var request []PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, _ = range request {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		sleepDuration := time.Duration(r.Intn(3-1+1)+1) * time.Second
		time.Sleep(sleepDuration)
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "결제 완료",
	})
}
