package discountpolicy

import (
	"math/big"
	"sort"
)

type CardDiscountPolicyByDollar struct {
	Threshold      *big.Float
	DiscountAmount *big.Float
}

var CardDiscountPolicies = []CardDiscountPolicyByDollar{
	{Threshold: big.NewFloat(50), DiscountAmount: big.NewFloat(8)},
	{Threshold: big.NewFloat(100), DiscountAmount: big.NewFloat(16)},
	{Threshold: big.NewFloat(150), DiscountAmount: big.NewFloat(30)},
	{Threshold: big.NewFloat(300), DiscountAmount: big.NewFloat(60)},
}

func GetCardPoliciesByDiscountAsc() []CardDiscountPolicyByDollar {
	policies := make([]CardDiscountPolicyByDollar, len(CardDiscountPolicies))
	copy(policies, CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCardPoliciesByDiscountDesc() []CardDiscountPolicyByDollar {
	policies := make([]CardDiscountPolicyByDollar, len(CardDiscountPolicies))
	copy(policies, CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}
