package handlers

import (
	"golang-server/internal/dto"
	"golang-server/internal/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	var request dto.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.Payment(request)

	c.JSON(http.StatusOK, dto.Success(nil))
}

func PaymentAll(c *gin.Context) {
	var request []dto.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	for _, req := range request {
		wg.Add(1)
		go func(r dto.PaymentRequest) {
			defer wg.Done()
			services.Payment(r)
		}(req)

	}
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"content": "결제 완료",
	})
}
