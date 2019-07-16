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
func LongestPalindrome(s string) string {
	switch METHOD {
	case METHOD_MINE:
		return longestPalindromeMine(s)
	case METHOD_LCA:
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

/**
思路：动态规划
给出定义　P(i,j) = true, 如果子串s[i]...s[j]是回文子串
              false, 其他情况
因此，
P(i,j) = (P(i+1,j-1) and s[i] == s[j])
基本示例如下：
P(i,i) = true
p(i,i+1) = (s[i] == s[i+1])
这就产生了一个直观的动态规划解法，我们首先初始化一字母和二字母的回文，然后找到所有三字母回文，并以此类推

复杂度分析：
时间复杂度: O(n^2)
空间复杂度: O(n^2)
 */
func longestPalindromeLCA(s string) string {
	slen := len(s)
	if slen == 0 {
		return ""
	}
	sbytes := []byte(s)
	lps := string(sbytes[0:1])
	//定义二维数组
	var pstore [][]bool
	for i := 0; i < slen; i++ {
		//定义变量或make生成变量后go会填充默认值 false
		pstore = append(pstore, make([]bool, slen))
	}
	//先初始化所有一个字母和二个字母的回文
	//然后计算所有回文并存储和使用中间结果，这样就将验证是否是回文这个操作的时间复杂度降低到了O(1)，整体时间复杂度也就降低到了O(n^2)
	//由于i依赖于i+1的结果，所以我们使用倒序来计算
	for i := slen-1; i >= 0; i-- {
		for j := 0; j < slen-i; j++ {
			if j <= 1 {
				pstore[i][i+j] = sbytes[i] == sbytes[i+j]
			} else {
				pstore[i][i+j] = pstore[i+1][i+j-1] && sbytes[i] == sbytes[i+j]
			}
			if pstore[i][i+j] && len(lps) <= j+1 {
				lps = string(sbytes[i:i+j+1])
			}
		}
	}
	return lps
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
