package discountpolicy

import (
	"math/big"
)

type DiscountPolicyInDollar struct {
	Threshold      *big.Float
	DiscountAmount *big.Float
}
