package problems

import (
	"bytes"
	"fmt"
	"utils"
)

/**
题目描述：
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
func LengthOfLongestSubstring(s string, method string) int {
	switch method {
	case "mine":
		return lengthOfLongestSubstringMine(s)
	case "lca":
		return lengthOfLongestSubstringLCA(s)
	}
	return 0
}

/**
思路：遍历字符串，枚举出所有的子串，然后计算长度，找出最长的子串

复杂度分析：
时间复杂度：O(n^3) n 为字符串长度
空间复杂度：O(min(n,m)) m为最长子串的长度
 */
func lengthOfLongestSubstringMine(s string) int {
	tempmap := make(map[string]int)
	length := 0
	sbytes := utils.StringBytes(s)
	for i := 0; i < len(sbytes); i++ {
		tempbytes := make([]byte, 0)
		tempbytes = append(tempbytes, sbytes[i])
		for j := i+1; j < len(sbytes); j++ {
			index := bytes.IndexByte(tempbytes, sbytes[j])
			if index == -1 {
				tempbytes = append(tempbytes, sbytes[j])
			}
			if index != -1 || j == len(sbytes)-1 {
				if len(tempbytes) > length {
					length = len(tempbytes)
				}
				tempmap[string(tempbytes)] = len(tempbytes)
				break
			}
		}
	}
	fmt.Println(tempmap)
	return length
}

/**
思路：滑动窗口
滑动窗口是数组/字符串问题中常用的抽象概念。 窗口通常是在数组/字符串中由开始和结束索引定义的一系列元素的集合，即 [i,j)（左闭，右开）。
而滑动窗口是可以将两个边界向某一方向“滑动”的窗口。例如，我们将 [i,j) 向右滑动 1 个元素，则它将变为 [i+1,j+1)（左闭，右开）
假设字符集为 ASCII 128，我们可以使用 [128]int 数组来定义字符到索引的映射，如果 s[j] 在 [i,j) 范围内有与 j'重复的字符，
我们不需要逐渐增加 i 。 我们可以直接跳过 [i，j']范围内的所有元素，并将 i 变为 j' + 1

复杂度分析：
时间复杂度：O(n) n是字符串的长度
空间复杂度：O(m) m是字符集的大小，如ASCII 128字符集的大小为128
 */
func lengthOfLongestSubstringLCA(s string) int {
	var temp [128]int = [128]int{0}
	sbytes,n,max := []byte(s),len(s),0
	for i,j := 0,0; j < n; j++ {
		if temp[sbytes[j]] > i {
			i = temp[sbytes[j]]
		}
		if j-i+1 > max {
			max = j-i+1
		}
		temp[sbytes[j]] = j+1
	}

	return max
}