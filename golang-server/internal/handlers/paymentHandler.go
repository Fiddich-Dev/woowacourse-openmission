package handlers

import (
	"golang-server/internal/dto"
	"golang-server/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	var request dto.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.Payment(request)

	c.JSON(http.StatusOK, gin.H{
		"content": request.GoodsList,
	})
}

//func PaymentAll(c *gin.Context) {
//	var request []dto.PaymentRequest
//
//	if err := c.ShouldBindJSON(&request); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	for _, _ = range request {
//		r := rand.New(rand.NewSource(time.Now().UnixNano()))
//		sleepDuration := time.Duration(r.Intn(3-1+1)+1) * time.Second
//		time.Sleep(sleepDuration)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"content": "결제 완료",
//	})
//}
