package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	length_max := 0
	var tmp string
	for i_s, v_s := range s {
		tmp = tmp + string(v_s)
		// fmt.Println("Temp: ", tmp)
		if i_s == len(s)-1 {
			break
		}
		for i_tmp, v_tmp := range tmp {
			target := s[i_s+1 : i_s+2]
			for _, v_target := range target { // convert to rune
				// fmt.Println("origin:", string(v_tmp), ",target: ", string(v_target))
				if v_tmp == v_target {
					// fmt.Println("Matched: Length is ", len(tmp))
					if length_max < len(tmp) {
						length_max = len(tmp)
					}
					tmp = tmp[i_tmp+1:]
					break
				}
			}
		}
	}

	if length_max < len(tmp) {
		length_max = len(tmp)
	}
	return length_max
}

func main() {
	// Input: s = "abcabcbb"
	// Output: 3
	// Explanation: The answer is "abc", with the length of 3.
	s := "abcabcbb"
	// s := "pwwkew"
	o := lengthOfLongestSubstring(s)
	fmt.Println(o)
}
