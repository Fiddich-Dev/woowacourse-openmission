package handlers

import (
	"golang-server/internal/dto"
	"golang-server/internal/services"
	"golang-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCouponDiscountPolicy(c *gin.Context) {
	couponDiscountPolicies := services.GetCouponPoliciesByDiscountAsc()

	response := make([]dto.DiscountPolicyByResponse, 0, len(couponDiscountPolicies))
	for _, policy := range couponDiscountPolicies {
		response = append(response, dto.DiscountPolicyByResponse{
			utils.BigFloatToString(policy.Threshold),
			utils.BigFloatToString(policy.DiscountAmount),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

func GetCardDiscountPolicy(c *gin.Context) {
	cardDiscountPolicies := services.GetCardPoliciesByDiscountAsc()

	response := make([]dto.DiscountPolicyByResponse, 0, len(cardDiscountPolicies))
	for _, policy := range cardDiscountPolicies {
		response = append(response, dto.DiscountPolicyByResponse{
			utils.BigFloatToString(policy.Threshold),
			utils.BigFloatToString(policy.DiscountAmount),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}
