package problems

/**
题目描述：
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
 */
func LongestPalindrome(s string, method string) string {
	switch method {
	case "mine":
		return longestPalindromeMine(s)
	case "lca":
		return longestPalindromeLCA(s)
	}

	return ""
}

/**
思路：遍历字符串的所有子串(第一个字符和最后一个字符相同的子串)，如果是回文则判断长度并保存

复杂度分析：
时间复杂度：O(n^3) 遍历子串为O(n^2) 检验是否回文为 O(n)
空间复杂度：O(n)
 */
func longestPalindromeMine(s string) string {
	if len(s) == 0 {
		return ""
	}
	sbytes := []byte(s)
	slen := len(sbytes)
	lps := string(sbytes[0:1])
	for i := 0; i < slen; i++ {
		for j := i+1; j < slen; j++ {
			//只有首尾字符相同时才有可能是回文
			if sbytes[i] == sbytes[j] && isPalindrome(sbytes[i:j+1]){
				if len(lps) < (j-i+1) {
					lps = string(sbytes[i:j+1])
				}
			}
		}
	}
	return lps
}

func longestPalindromeLCA(s string) string {
	return ""
}

func isPalindrome(sbytes []byte) bool {
	slen := len(sbytes)
	for i := 0; i < slen/2; i++ {
		if sbytes[i] != sbytes[slen-1-i] {
			return false
		}
	}
	return true
}
