package main

import (
	"math/big"
	"sort"
)

type CouponDiscountPolicyByDollar struct {
	Threshold      *big.Float
	DiscountAmount *big.Float
}

var CouponDiscountPolicies = []CouponDiscountPolicyByDollar{
	{Threshold: big.NewFloat(25), DiscountAmount: big.NewFloat(4)},
	{Threshold: big.NewFloat(54), DiscountAmount: big.NewFloat(8)},
	{Threshold: big.NewFloat(89), DiscountAmount: big.NewFloat(13)},
	{Threshold: big.NewFloat(209), DiscountAmount: big.NewFloat(31)},
	{Threshold: big.NewFloat(679), DiscountAmount: big.NewFloat(102)},
}

func GetCouponPoliciesByDiscountAsc() []CouponDiscountPolicyByDollar {
	policies := make([]CouponDiscountPolicyByDollar, len(CouponDiscountPolicies))
	copy(policies, CouponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCouponPoliciesByDiscountDesc() []CouponDiscountPolicyByDollar {
	policies := make([]CouponDiscountPolicyByDollar, len(CouponDiscountPolicies))
	copy(policies, CouponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}
