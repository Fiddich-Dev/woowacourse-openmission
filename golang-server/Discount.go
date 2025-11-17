package main

import "math/big"

func applyCardDiscount(price *big.Float) *big.Float {
	policies := GetCardPoliciesByDiscountDesc()

	for _, p := range policies {
		if price.Cmp(p.Threshold) >= 0 {
			result := new(big.Float).Sub(price, p.DiscountAmount)
			return result
		}
	}

	return price
}

func applyCouponDiscount(price *big.Float) *big.Float {
	policies := GetCouponPoliciesByDiscountDesc()

	for _, p := range policies {
		if price.Cmp(p.Threshold) >= 0 {
			result := new(big.Float).Sub(price, p.DiscountAmount)
			return result
		}
	}

	return price
}

func ApplyDiscount(price *big.Float) *big.Float {
	return applyCardDiscount(applyCouponDiscount(price))
}
