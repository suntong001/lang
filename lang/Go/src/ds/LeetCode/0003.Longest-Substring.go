package main

import "fmt"

/*

 Given a string s, find the length of the longest substring
 without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.


Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


*/

func main() {
	lengthOfLongestSubstring := lengthOfLongestSubstring0
	lengthOfLongestSubstring = lengthOfLongestSubstring1
	fmt.Println(lengthOfLongestSubstring("abcbbcbb"))
	fmt.Println("---")
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println("---")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println("---")
}

func lengthOfLongestSubstring1(s string) int {
	ll, rr, res, str := 0, 0, 0, ""
	lastAt := make(map[byte]int, len(s)) // last pos such char was at

	for rr < len(s) {
		if idx, ok := lastAt[s[rr]]; ok && idx >= ll {
			ll = idx + 1 // need to be at next pos to last-seen to avoid repeating
		}

		lastAt[s[rr]] = rr
		rr++
		//res = max(res, rr-ll)
		if res < rr-ll {
			res = rr - ll
			str = s[ll:rr]
		}
		fmt.Println("  ", ll, rr, str, res, lastAt)
	}

	return res
}

func lengthOfLongestSubstring0(s string) int {
	res, left, right := 0, 0, -1
	if res < right-left {
		res = right - left
	}
	return res
}
