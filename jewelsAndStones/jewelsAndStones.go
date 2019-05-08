package leetcode

func numJewelsInStones(J string, S string) int {
	result := 0
	seen := make(map[rune]bool, 0)
	for _, l := range J {
		seen[l] = true
	}

	for _, l := range S {
		_, ok := seen[l]
		if ok {
			result++
		}
	}

	return result
}
