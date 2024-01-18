package calculator

import (
	"math"
)

//go:generate mockgen -source=calculator.go -destination=./mocks/packsGetAller.go -package=mocks
type PacksGetAller interface {
	GetAll() []int
}

type PacksCalculator struct {
	provider PacksGetAller
}

func New(provider PacksGetAller) *PacksCalculator {
	return &PacksCalculator{
		provider: provider,
	}
}

func (c *PacksCalculator) Calculate(order int) []int {
	packSizes := c.provider.GetAll()
	return calculateOptimalPacks(order, packSizes)
}

func calculateOptimalPacks(order int, packSizes []int) []int {
	if len(packSizes) == 0 {
		return []int{}
	}

	target := order + max(packSizes)

	n := len(packSizes)
	table := make([]int, target+1)
	for i := 1; i <= target; i++ {
		table[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		packSize := packSizes[i]
		for quantity := packSize; quantity <= target; quantity++ {
			table[quantity] = min(table[quantity], table[quantity-packSize]+1)
		}
	}

	smallestTarget := order
	for i := order; i < len(table); i++ {
		if table[i] != math.MaxInt32 {
			smallestTarget = i
			break
		}
	}

	remainingQuantity := smallestTarget
	var packs []int
	for i := 0; i < n && remainingQuantity > 0; i++ {
		packSize := packSizes[i]
		for remainingQuantity >= packSize && table[remainingQuantity-packSize]+1 == table[remainingQuantity] {
			packs = append(packs, packSize)
			remainingQuantity -= packSize
		}
	}

	return packs
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(s []int) int {
	if len(s) < 1 {
		return -1
	}

	max := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}
