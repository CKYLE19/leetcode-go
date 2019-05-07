package leetcode

func smallestChar(letters []byte, target byte) byte {
	wrapAround := letters[0]
	for i := 0; i < len(letters); i++ {
		if target < letters[i] {
			return letters[i]
		}
	}
	return wrapAround
}