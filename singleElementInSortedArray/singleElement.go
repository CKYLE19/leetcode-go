package leetcode

func singleElement(arr []int) int {
	seen := make(map[int]int, 0)
	for _, num := range arr {
		_, ok := seen[num]
		if ok {
			seen[num]++
		} else {
			seen[num] = 1
		}
	}

	for k, v := range seen {
		if v == 1 {
			return k
		}
	}
	return -1
}
