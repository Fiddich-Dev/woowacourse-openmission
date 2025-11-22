package handlers

import (
	"golang-server/internal/config"
	"golang-server/internal/dto"
	"golang-server/internal/services"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	var request dto.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, config.RequestErrorMessage))
		return
	}

	services.Payment(request)

	c.JSON(http.StatusOK, dto.Success(nil))
}

func PaymentAll(c *gin.Context) {
	var request []dto.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail(http.StatusBadRequest, config.RequestErrorMessage))
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

	c.JSON(http.StatusOK, dto.Success(nil))
}
