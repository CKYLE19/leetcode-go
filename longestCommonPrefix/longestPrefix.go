package leetcode

func longestPrefix(words []string) string {
	result := ""
	if len(words) == 0 {
		return result
	}
	i := 0
	smallest := 0
	for j, w := range words {
		if len(w) < len(words[smallest]) {
			smallest = j
		}
	}
	for true {
		if i == len(words[smallest]) {
			return result
		}
		c := words[smallest][i]
		for _, word := range words {
			if word[i] != c {
				return result
			}
		}
		result = result + string(c)
		i++
	}
	return result
}
