package longestSubstringWithoutRepeatingChars

func longestSubstring(runeStr []rune, begin int, end int) string {
	if begin < 0 || end >= len(runeStr) || begin > end {
		return ""
	}
	return string(runeStr[begin : end+1])
}

func Solution(s string) int {
	runeStr := []rune(s)
	window := make(map[rune]bool)
	begin := 0
	end := begin
	low := 0
	high := low

	for high < len(s) {
		if window[runeStr[high]] {
			for string(runeStr[low]) != string(runeStr[high]) {
				window[runeStr[low]] = false
				low += 1
			}
			low += 1
		} else {
			window[runeStr[high]] = true

			if end-begin < high-low {
				begin = low
				end = high
			}
		}
		high += 1
	}
	return len(longestSubstring(runeStr, begin, end))
}
