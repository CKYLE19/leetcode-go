package leetcode

func twoSum(nums []int, target int) []int {
	compliment := make(map[int]int, 0)
	for i, num := range nums {
		compliment[target-num] = i
	}
	for i, num := range nums {
		if j, ok := compliment[num]; ok && i != j {
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
