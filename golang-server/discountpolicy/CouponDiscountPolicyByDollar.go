package discountpolicy

import (
	"math/big"
	"sort"
)

type CouponDiscountPolicyInDollar struct {
	DiscountPolicyInDollar
}

func newCouponDiscountPolicyInDollar(threshold, discountAmount *big.Float) CouponDiscountPolicyInDollar {
	return CouponDiscountPolicyInDollar{
		DiscountPolicyInDollar{
			threshold,
			discountAmount,
		},
	}
}

var couponDiscountPolicies = []CouponDiscountPolicyInDollar{
	newCouponDiscountPolicyInDollar(big.NewFloat(25), big.NewFloat(4)),
	newCouponDiscountPolicyInDollar(big.NewFloat(54), big.NewFloat(8)),
	newCouponDiscountPolicyInDollar(big.NewFloat(89), big.NewFloat(13)),
	newCouponDiscountPolicyInDollar(big.NewFloat(209), big.NewFloat(31)),
	newCouponDiscountPolicyInDollar(big.NewFloat(679), big.NewFloat(102)),
}

func GetCouponPoliciesByDiscountAsc() []CouponDiscountPolicyInDollar {
	policies := make([]CouponDiscountPolicyInDollar, len(couponDiscountPolicies))
	copy(policies, couponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCouponPoliciesByDiscountDesc() []CouponDiscountPolicyInDollar {
	policies := make([]CouponDiscountPolicyInDollar, len(couponDiscountPolicies))
	copy(policies, couponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}
