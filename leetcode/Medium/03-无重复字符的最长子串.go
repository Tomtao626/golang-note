package main

import "fmt"

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	cnt := [128]int{}
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var s = "abcaadagaeadb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
