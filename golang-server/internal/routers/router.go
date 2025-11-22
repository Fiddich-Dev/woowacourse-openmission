package routers

import (
	"golang-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	goods := router.Group("/goods")
	{
		goods.POST("/partition", handlers.GoodsListPartition)
	}

	discountPolicy := router.Group("/discount-policy")
	{
		discountPolicy.GET("/coupon", handlers.GetCouponDiscountPolicy)
		discountPolicy.GET("/card", handlers.GetCardDiscountPolicy)
	}

	exchangeRate := router.Group("/exchange-rate")
	{
		exchangeRate.GET("", handlers.GetUsdRateHandler)
	}

	payment := router.Group("/payment")
	{
		payment.POST("", handlers.Payment)
		payment.POST("all", handlers.PaymentAll)
	}

	return router
}
