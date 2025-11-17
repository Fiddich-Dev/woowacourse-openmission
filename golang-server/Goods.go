package main

import (
	"encoding/json"
	"math/big"
)

type Goods struct {
	Name  string     `json:"name"`
	Price *big.Float `json:"price"`
}

func (g *Goods) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	g.Name = aux.Name

	f, _, err := big.ParseFloat(aux.Price, 10, 256, big.ToNearestEven)
	if err != nil {
		return err
	}
	g.Price = f

	return nil
}
