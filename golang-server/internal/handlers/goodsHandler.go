package handlers

import (
	"golang-server/internal/dto"
	"golang-server/internal/models"
	"golang-server/internal/services"
	"math/big"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func GoodsListPartition(c *gin.Context) {
	var request dto.GoodsListPartitionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	partitionedGoods := services.Partition(request.GoodsList)
	response := convertPartitionedGoodsToResponse(partitionedGoods)

	c.JSON(http.StatusOK, dto.Success(response))
}

func convertPartitionedGoodsToResponse(partitions [][][]models.Goods) []dto.GoodsListPartitionResponse {
	beforeTotalPrice := big.NewFloat(0)
	for _, goodsList := range partitions[0] {
		beforeTotalPrice = beforeTotalPrice.Add(beforeTotalPrice, calcBeforePrice(goodsList))
	}

	bestPartition := findBestPartition(partitions)

	var result []dto.GoodsListPartitionResponse
	for _, goods := range bestPartition {
		beforePrice := calcBeforePrice(goods)
		afterPrice := calcAfterPrice(goods)

		result = append(result, dto.GoodsListPartitionResponse{
			goods,
			beforePrice.String(),
			afterPrice.String()})
	}

	return result
}

func findBestPartition(partitions [][][]models.Goods) [][]models.Goods {
	sort.Slice(partitions, func(i, j int) bool {
		totalI := calcAfterPriceForPartition(partitions[i])
		totalJ := calcAfterPriceForPartition(partitions[j])
		return totalI.Cmp(totalJ) < 0
	})

	return partitions[0]
}

func calcAfterPriceForPartition(groups [][]models.Goods) *big.Float {
	total := big.NewFloat(0)
	for _, group := range groups {
		groupAfter := calcAfterPrice(group)
		total.Add(total, groupAfter)
	}
	return total
}

func calcBeforePrice(goods []models.Goods) *big.Float {
	sum := big.NewFloat(0)

	for _, g := range goods {
		sum.Add(sum, g.Price)
	}
	return sum
}

func calcAfterPrice(goods []models.Goods) *big.Float {
	beforePrice := calcBeforePrice(goods)
	return services.Discount(beforePrice)
}
