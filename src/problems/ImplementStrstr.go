package problems

import "utils"

/**
题目描述：

实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 */
func StrStr(haystack string, needle string) int {
	switch METHOD {
	case METHOD_MINE:
		return strStrMine(haystack, needle)
	case METHOD_LCA:
		return strStrLCA(haystack, needle)
	}

	return -1
}

/**
思路：
匹配遍历字符串逐个字符做比较，如果有字符和needle中的第一个字符相同则遍历needle继续比较剩余的字符，直到结束

算法：
可在第一个for循环中根据剩余长度做优化，如果剩余长度小于needle长度则不用比较直接返回-1，
比较needle中剩余字符时如果遇到不匹配的字符则需要返回第一层循环继续遍历，不能直接return -1，因为后面可能还有相同的字符串

复杂度分析：
时间复杂度：O((m-n)*n) = O(n^2) 假设 haystack长度为m，needle长度为n
时间复杂度：O(1) 在使用了字符串原有内存转字符数组的情况下不额外占用空间
 */
func strStrMine(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	start := -1
	haystackbytes := utils.StringBytes(haystack)
	lenhay := len(haystackbytes)
	needlebytes := utils.StringBytes(needle)
	lennee := len(needlebytes)
	for i := 0; i < lenhay; i++  {
		if i+lennee > lenhay {
			return -1
		}
		if haystackbytes[i] == needlebytes[0] {
			start = i
			for j := 1; j < lennee; j++ {
				k := i+j
				if k >= lenhay || haystackbytes[k] != needlebytes[j] {
					start = -1
					break
				}
			}
			if start != -1 {
				return start
			}
		}
	}
	return -1
}

//LCA暂无题解
func strStrLCA(haystack string, needle string) int {
	return -1
}
