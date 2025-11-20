package repositories

import (
	"golang-server/internal/models"
	"math/big"
)

var CouponDiscountPolicies = []models.DiscountPolicyInDollar{
	{big.NewFloat(25), big.NewFloat(4)},
	{big.NewFloat(54), big.NewFloat(8)},
	{big.NewFloat(89), big.NewFloat(13)},
	{big.NewFloat(209), big.NewFloat(31)},
	{big.NewFloat(679), big.NewFloat(102)},
}

var CardDiscountPolicies = []models.DiscountPolicyInDollar{
	{big.NewFloat(50), big.NewFloat(8)},
	{big.NewFloat(100), big.NewFloat(16)},
	{big.NewFloat(150), big.NewFloat(30)},
	{big.NewFloat(300), big.NewFloat(60)},
}
