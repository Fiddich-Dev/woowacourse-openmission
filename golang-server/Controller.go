package main

import (
	"golang-server/discount"
	"golang-server/model"
	"math/big"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type PartitionRequest struct {
	Goods []model.Goods `json:"goods"`
}

type PartitionResponse struct {
	Goods       []model.Goods `json:"goods"`
	BeforePrice string        `json:"before_price"`
	AfterPrice  string        `json:"after_price"`
}

func PartitionService(c *gin.Context) {
	var request PartitionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	partitionedGoods := Partition(request.Goods)

	response := ConvertToResponse(partitionedGoods)

	c.JSON(http.StatusOK, gin.H{
		"content": response,
	})
}

func calcBeforePrice(goods []model.Goods) *big.Float {
	sum := big.NewFloat(0)

	for _, g := range goods {
		sum.Add(sum, g.Price)
	}
	return sum
}

func calcAfterPrice(goods []model.Goods) *big.Float {
	beforePrice := calcBeforePrice(goods)
	return discount.ApplyDiscount(beforePrice)
}

func ConvertToResponse(partitions [][][]model.Goods) []PartitionResponse {
	beforeTotalPrice := big.NewFloat(0)
	for _, goodsList := range partitions[0] {
		beforeTotalPrice = beforeTotalPrice.Add(beforeTotalPrice, calcBeforePrice(goodsList))
	}

	sort.Slice(partitions, func(i, j int) bool {
		totalI := calcAfterPriceForPartition(partitions[i])
		totalJ := calcAfterPriceForPartition(partitions[j])
		return totalI.Cmp(totalJ) < 0
	})

	bestPartition := partitions[0]

	var result []PartitionResponse

	for _, goods := range bestPartition {
		beforePrice := calcBeforePrice(goods)
		afterPrice := calcAfterPrice(goods)

		result = append(result, PartitionResponse{goods, beforePrice.String(), afterPrice.String()})
	}

	return result
}

func calcAfterPriceForPartition(groups [][]model.Goods) *big.Float {
	total := big.NewFloat(0)
	for _, group := range groups {
		groupAfter := calcAfterPrice(group)
		total.Add(total, groupAfter)
	}
	return total
}
