package problems

/**
题目描述：
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。



示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"

https://leetcode-cn.com/problems/count-and-say/
 */
func CountAndSay(n int, method string) string {
	switch method {
	case "mine":
		return countAndSayMine(n)
	case "lca":
		return countAndSayLCA(n)
	}

	return ""
}

/**
递归调用 f(n)->f(n-1)...f(1)
复杂度分析：
时间复杂度：O(n*lgn) n为递归深度
空间复杂度：
 */
func countAndSayMine(n int) string {
	if n == 1 {
		return "1"
	} else {
		str := countAndSayMine(n-1)
		var result []byte
		strbytes := []byte(str)
		lenstrbytes := len(strbytes)
		var last byte = strbytes[0]
		i := 0
		for j := 0; j < lenstrbytes; j++ {
			if last == strbytes[j] {
				i++
			} else {
				result = append(result, byte('0'+i), last)
				last = strbytes[j]
				i = 1
			}

			if j == lenstrbytes-1 {
				result = append(result, byte('0'+i), last)
			}
		}
		return string(result)
	}
}

/**
LCA 暂无解答
 */
func countAndSayLCA(n int) string {
	return ""
}
