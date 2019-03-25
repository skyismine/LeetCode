package problems

import (
	"commonds"
	"math"
)

/**
题目描述:
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z

https://leetcode-cn.com/problems/longest-common-prefix/
 */
func LongestCommonPrefix(strs []string, method string) string {
	switch method {
	case "mine":
		return longestCommonPrefixMine(strs)
	case "lca":
		return longestCommonPrefixLCA(strs)
	}

	return ""
}

/**
 先计算字符串数组中最短长度minlen，然后遍历minlen次，对字符串数组中每个字符串比较第i个字符是否相同，如果不同则返回

算法效率：
   时间复杂度：O(minlen*n) minlen为数组中最短字符串的长度，n为字符串数组的长度
   空间复杂度: O(len(strs[0])+len(strs[1])+...+len(strs[n]))
 */
func longestCommonPrefixMine(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	var minlen int = math.MaxInt32
	for _, value := range strs {
		if minlen > len(value) {
			minlen = len(value)
		}
	}

	if minlen == 0 {
		return ""
	}

	var ret []byte
	var firstbytes []byte = []byte(strs[0])

	for i := 0; i < minlen; i++ {
		for j := 1; j < len(strs); j++ {
			curr := []byte(strs[j])
			if firstbytes[i] != curr[i] {
				return string(ret)
			}
		}
		ret = append(ret, firstbytes[i])
	}
	return string(ret)
}

/**
 我们可以通过将所有的键值 S 存储到一颗字典树中来优化最长公共前缀查询操作。 如果你想获得更多关于字典树的信息，可以查看这篇文章 Implement a trie (Prefix trie)[https://leetcode.com/articles/implement-trie-prefix-tree/] 。在字典树中，从根向下的每一个节点都代表一些键值的公共前缀。 但是我们需要找到字符串q 和所有键值字符串的最长公共前缀。 这意味着我们需要从根找到一条最深的路径，满足以下条件：

这是所查询的字符串 q 的一个前缀

路径上的每一个节点都有且仅有一个孩子。 否则，找到的路径就不是所有字符串的公共前缀

路径不包含被标记成某一个键值字符串结尾的节点。 因为最长公共前缀不可能比某个字符串本身长

算法

最后的问题就是如何找到字典树中满足上述所有要求的最深节点。 最有效的方法就是建立一颗包含字符串 [S1...Sn] 的字典树。 然后在这颗树中匹配 q 的前缀。 我们从根节点遍历这颗字典树，直到因为不能满足某个条件而不能再遍历为止。

复杂度分析

最坏情况下查询字符串 q 的长度为 m 并且它与数组中 n 个字符串均相同。

时间复杂度：预处理过程 O(S)，其中 S 是数组里所有字符串中字符数量的总和，最长公共前缀查询操作的复杂度为 O(m)。

建立字典树的时间复杂度为 O(S)。 在字典树中查找字符串 q 的最长公共前缀在最坏情况下需要 O(m) 的时间。

空间复杂度：O(S)， 我们只需要使用额外的 S 空间建立字典树。
 */
func longestCommonPrefixLCA(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	root := commonds.NewTireRoot()
	for _,value := range strs {
		if len(value) == 0 {
			return ""
		}
		root.Insert(value)
	}
	return root.SearchLCP(strs[0])
}
