package leetcode

func longestSubstring(s string) string {
	left, right := 0, 0
	seen := make(map[string]int, 0)
	var longest string
	for right < len(s) {
		if _, ok := seen[string(s[right])]; ok && seen[string(s[right])] >= left {
			// This is how we know we have seen within bounds
			// In this case we set to seen[right]+1
			left = seen[string(s[right])] + 1
		}
		seen[string(s[right])] = right
		if right-left+1 > len(longest) {
			longest = s[left : right+1]
		}
		right++
	}
	return longest
}
