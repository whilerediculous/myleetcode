package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test42_TrappingRainWater(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := TrappingRainWaterIterate(nums)
	assert.Equal(t, 6, ret)
}
