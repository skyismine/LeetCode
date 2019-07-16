package problems

/**
 题目描述：
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

https://leetcode-cn.com/problems/roman-to-integer/
 */
func RomanToInt(s string) int {
	switch METHOD {
	case METHOD_MINE:
		return romanToIntMine(s)
	case METHOD_MINE2:
		return romanToIntMine2(s)
	case METHOD_LCA:
		return romanToIntLCA(s)
	}

	return 0
}

/**
 遍历字符串对于I X C字符的处理需要结合它的下一个字符来进行
 复杂度分析：
     时间复杂度为O(n)
     空间复杂度为O(n)
 */
func romanToIntMine(s string) int {
	strbytes := []byte(s)
	strlen := len(strbytes)
	ret := 0
	for i := 0; i < strlen; i++ {
		switch strbytes[i] {
		case 'I':
			if i >= strlen-1 {
				ret += 1
			} else {
				if strbytes[i+1] == 'V' {
					ret += 4
					i++
				} else if (strbytes[i+1] == 'X') {
					ret += 9
					i++
				} else {
					ret += 1
				}
			}
		case 'V':
			ret += 5
		case 'X':
			if i >= strlen-1 {
				ret += 10
			} else {
				if strbytes[i+1] == 'L' {
					ret += 40
					i++
				} else if (strbytes[i+1] == 'C') {
					ret += 90
					i++
				} else {
					ret += 10
				}
			}
		case 'L':
			ret += 50
		case 'C':
			if i >= strlen-1 {
				ret += 100
			} else {
				if strbytes[i+1] == 'D' {
					ret += 400
					i++
				} else if (strbytes[i+1] == 'M') {
					ret += 900
					i++
				} else {
					ret += 100
				}
			}
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}
	return ret
}

/**
 将可能的值都存储在map中，利用I*100+V X*100+L等几个可能的值不相等不会导致key冲突的特性使用int作为键，遍历字符串如果遇到I X C则需要结合下一个字符
共同处理，否则直接加当前处理结果
 复杂度分析：
     时间复杂度为O(n)
     空间复杂度为O(n)
  没有 romanToIntMine 效率高，经打印耗时分析 go map 查找效率很低，一次查找在 300ns 左右，且速度不稳定，还有500ns以上的
 */
func romanToIntMine2(s string) int {
	ixcmap := map[int]int{
		'I' : 1,
		'I'*100+'V' : 4,
		'I'*100+'X' : 9,
		'V' : 5,
		'X' : 10,
		'X'*100+'L' : 40,
		'X'*100+'C' : 90,
		'L' : 50,
		'C' : 100,
		'C'*100+'D' : 400,
		'C'*100+'M' : 900,
		'D' : 500,
		'M' : 1000,
	}
	strbytes := []byte(s)
	strlen := len(strbytes)
	ret := 0
	for i := 0; i < strlen; i++ {
		var curr int = int(strbytes[i])
		if i < strlen-1 && (curr == 'I' || curr == 'X' || curr == 'C') {
			var next int = int(strbytes[i+1])
			key := curr*100+next
			if v, ok := ixcmap[key]; ok {
				ret += v
				i++
				continue
			}
		}

		ret += ixcmap[curr]
	}
	return ret
}

/**
 LCA 上没有解答
 */
func romanToIntLCA(s string) int {
	return 0
}
