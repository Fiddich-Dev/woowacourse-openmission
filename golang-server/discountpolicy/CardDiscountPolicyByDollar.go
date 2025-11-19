package discountpolicy

import (
	"math/big"
	"sort"
)

type CardDiscountPolicyInDollar struct {
	DiscountPolicyInDollar
}

func newCardDiscountPolicyInDollar(threshold, discountAmount *big.Float) CardDiscountPolicyInDollar {
	return CardDiscountPolicyInDollar{
		DiscountPolicyInDollar{
			threshold,
			discountAmount,
		},
	}
}

var CardDiscountPolicies = []CardDiscountPolicyInDollar{
	newCardDiscountPolicyInDollar(big.NewFloat(50), big.NewFloat(8)),
	newCardDiscountPolicyInDollar(big.NewFloat(100), big.NewFloat(16)),
	newCardDiscountPolicyInDollar(big.NewFloat(150), big.NewFloat(30)),
	newCardDiscountPolicyInDollar(big.NewFloat(300), big.NewFloat(60)),
}

func GetCardPoliciesByDiscountAsc() []CardDiscountPolicyInDollar {
	policies := make([]CardDiscountPolicyInDollar, len(CardDiscountPolicies))
	copy(policies, CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCardPoliciesByDiscountDesc() []CardDiscountPolicyInDollar {
	policies := make([]CardDiscountPolicyInDollar, len(CardDiscountPolicies))
	copy(policies, CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}
