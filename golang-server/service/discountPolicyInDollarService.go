package service

import (
	"golang-server/discountpolicy"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiscountPolicyByDollarDto struct {
	Threshold      string `json:"threshold"`
	DiscountAmount string `json:"discountAmount"`
}

func CouponDiscountPolicyService(c *gin.Context) {
	couponDiscountPolicies := discountpolicy.CouponDiscountPolicies
	response := MapCouponPoliciesToDto(couponDiscountPolicies)

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

func CardDiscountPolicyService(c *gin.Context) {
	cardDiscountPolicies := discountpolicy.CardDiscountPolicies
	response := MapPoliciesToDto(cardDiscountPolicies)

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}
