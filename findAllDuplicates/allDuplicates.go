package leetcode

func findDuplicates(nums []int) []int {
	/**
	 * [4,3,2,7,8,2,3,1]
	 * => [2,3]
	 *
	 * do in constant space and O(n) time
	 *
	 * we can do in non constant space by keeping a seen set
	 * every time something is already in the set we say it's a duplicate
	 *
	 * we can do in constant space by swapping elements into their corresponding
	 * A[i]-1 index. if the space already contains that number, we have a duplicate
	 *
	 * need to note that if i == A[i]-1 when iterating, then we don't consider it a duplicate
	 */

	result := make([]int, 0)
	i := 0
	for i < len(nums) {
		if nums[i] < 1 {
			i++
			continue
		}
		if nums[i]-1 != i {
			if nums[i] == nums[nums[i]-1] {
				result = append(result, nums[i])
				nums[i] = 0
				i++
			} else {
				// we need to swap into correct position
				tmp := nums[nums[i]-1]
				nums[nums[i]-1] = nums[i]
				nums[i] = tmp
			}
		} else {
			i++
		}
	}

	return result
}
