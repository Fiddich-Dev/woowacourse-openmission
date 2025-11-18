package model

import (
	"golang-server/discountpolicy"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiscountPolicyByDollarDto struct {
	Threshold      string `json:"threshold"`
	DiscountAmount string `json:"discountAmount"`
}

func CardDiscountPolicyService(c *gin.Context) {
	cardDiscountPolicies := discountpolicy.CardDiscountPolicies
	response := MapPoliciesToDto(cardDiscountPolicies)

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

func CouponDiscountPolicyService(c *gin.Context) {
	couponDiscountPolicies := discountpolicy.CouponDiscountPolicies
	response := MapCouponPoliciesToDto(couponDiscountPolicies)

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

func MapPoliciesToDto(policies []discountpolicy.CardDiscountPolicyByDollar) []DiscountPolicyByDollarDto {
	result := make([]DiscountPolicyByDollarDto, 0, len(policies))

	for _, p := range policies {
		dto := DiscountPolicyByDollarDto{
			Threshold:      BigFloatToString(RoundTo2Decimal(p.Threshold)),
			DiscountAmount: BigFloatToString(RoundTo2Decimal(p.DiscountAmount)),
		}
		result = append(result, dto)
	}

	return result
}

func MapCouponPoliciesToDto(policies []discountpolicy.CouponDiscountPolicyByDollar) []DiscountPolicyByDollarDto {
	result := make([]DiscountPolicyByDollarDto, 0, len(policies))

	for _, p := range policies {
		dto := DiscountPolicyByDollarDto{
			Threshold:      BigFloatToString(RoundTo2Decimal(p.Threshold)),
			DiscountAmount: BigFloatToString(RoundTo2Decimal(p.DiscountAmount)),
		}
		result = append(result, dto)
	}

	return result
}

func BigFloatToString(f *big.Float) string {
	return f.Text('f', -1)
}

func RoundTo2Decimal(f *big.Float) *big.Float {
	m := big.NewFloat(100)

	scaled := new(big.Float).Mul(f, m)

	scaled.Add(scaled, big.NewFloat(0.5))

	floored, _ := scaled.Int(nil)

	result := new(big.Float).Quo(new(big.Float).SetInt(floored), m)

	return result
}
