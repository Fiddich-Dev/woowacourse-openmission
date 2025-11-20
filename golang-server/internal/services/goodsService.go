package services

import (
	"golang-server/internal/models"
	"math/big"
)

func Discount(price *big.Float) *big.Float {
	return cardDiscount(couponDiscount(price))
}

func cardDiscount(price *big.Float) *big.Float {
	policies := GetCardPoliciesByDiscountDesc()

	for _, p := range policies {
		if price.Cmp(p.Threshold) >= 0 {
			result := new(big.Float).Sub(price, p.DiscountAmount)
			return result
		}
	}

	return price
}

func couponDiscount(price *big.Float) *big.Float {
	policies := GetCouponPoliciesByDiscountDesc()

	for _, p := range policies {
		if price.Cmp(p.Threshold) >= 0 {
			result := new(big.Float).Sub(price, p.DiscountAmount)
			return result
		}
	}

	return price
}

func Partition(items []models.Goods) [][][]models.Goods {
	var result [][][]models.Goods
	backtracking(0, items, [][]models.Goods{}, &result)
	return result
}

func backtracking(index int, items []models.Goods, current [][]models.Goods, result *[][][]models.Goods) {
	if index == len(items) {
		copied := deepCopyGroups(current)
		*result = append(*result, copied)
		return
	}

	currentItem := items[index]

	for idx := range current {
		current[idx] = append(current[idx], currentItem)
		backtracking(index+1, items, current, result)
		current[idx] = current[idx][:len(current[idx])-1]
	}

	newGroup := []models.Goods{currentItem}
	current = append(current, newGroup)
	backtracking(index+1, items, current, result)
	current = current[:len(current)-1]
}

func deepCopyGroups(groups [][]models.Goods) [][]models.Goods {
	copyGroups := make([][]models.Goods, len(groups))
	for idx := range groups {
		copyGroups[idx] = make([]models.Goods, len(groups[idx]))
		copy(copyGroups[idx], groups[idx])
	}
	return copyGroups
}
