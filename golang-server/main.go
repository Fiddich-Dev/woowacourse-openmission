package main

import (
	"golang-server/internal/routers"
)

func main() {
	//// 할인정보 조회
	//r.GET("/coupon-discount-policy", discountpolicy.GetCouponDiscountPolicy)
	//r.GET("/card-discount-policy", discountpolicy.GetCardDiscountPolicy)
	//
	//// 상품들 분할
	//r.POST("/partition", service.GoodsListPartition)
	//
	//// 환율 조회
	//r.GET("/exchange-rate", exchangerate.GetExchangeRate)
	//
	//// 결제
	//r.POST("/payment", payment.Payment)

	app := routers.SetupRouter()

	app.Run()
}
