package services

import (
	"golang-server/internal/models"
	"golang-server/internal/repositories"
	"sort"
)

func GetCouponPoliciesByDiscountAsc() []models.DiscountPolicyInDollar {
	policies := make([]models.DiscountPolicyInDollar, len(repositories.CouponDiscountPolicies))
	copy(policies, repositories.CouponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCouponPoliciesByDiscountDesc() []models.DiscountPolicyInDollar {
	policies := make([]models.DiscountPolicyInDollar, len(repositories.CouponDiscountPolicies))
	copy(policies, repositories.CouponDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}

func GetCardPoliciesByDiscountAsc() []models.DiscountPolicyInDollar {
	policies := make([]models.DiscountPolicyInDollar, len(repositories.CardDiscountPolicies))
	copy(policies, repositories.CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) < 0
	})

	return policies
}

func GetCardPoliciesByDiscountDesc() []models.DiscountPolicyInDollar {
	policies := make([]models.DiscountPolicyInDollar, len(repositories.CardDiscountPolicies))
	copy(policies, repositories.CardDiscountPolicies)

	sort.Slice(policies, func(i, j int) bool {
		return policies[i].DiscountAmount.Cmp(policies[j].DiscountAmount) > 0
	})

	return policies
}
