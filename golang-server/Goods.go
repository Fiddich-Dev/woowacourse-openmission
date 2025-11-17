package main

import "math/big"

type Goods struct {
	Name  string     `json:"name"`
	Price *big.Float `json:"price"`
}
