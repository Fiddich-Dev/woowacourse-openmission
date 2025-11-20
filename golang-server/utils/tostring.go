package utils

import "math/big"

func BigFloatToString(f *big.Float) string {
	return f.Text('f', -1)
}
