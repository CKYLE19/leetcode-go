package leetcode

import "math"

func maxSubArray(nums []int) int {
	mem := make([]float64, len(nums))
	mem[len(nums)-1] = float64(nums[len(nums)-1])

	for i := len(nums) - 2; i >= 0; i-- {
		mem[i] = math.Max(float64(nums[i]), float64(nums[i])+mem[i+1])
	}

	result := math.Inf(-1)
	for _, num := range mem {
		result = math.Max(result, float64(num))
	}

	return int(result)
}
