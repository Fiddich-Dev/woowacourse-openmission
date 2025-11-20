package utils

import "math/big"

func RoundTo2Decimal(f *big.Float) *big.Float {
	m := big.NewFloat(100)

	scaled := new(big.Float).Mul(f, m)

	scaled.Add(scaled, big.NewFloat(0.5))

	floored, _ := scaled.Int(nil)

	result := new(big.Float).Quo(new(big.Float).SetInt(floored), m)

	return result
}
