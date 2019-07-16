package problems

import (
	"strings"
	"bytes"
)

/**
题目描述：
请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−2^31,  2^31 − 1]。如果数值超过这个范围，qing返回  INT_MAX (2^31 − 1) 或 INT_MIN (−2^31) 。

示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN (−231) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-integer-atoi
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func MyAtoi(str string) int {
	switch METHOD {
	case METHOD_MINE:
		return myAtoiMine(str)
	case METHOD_LCA:
		return myAtoiLCA(str)
	}

	return 0
}

/**
算法： step1: 规范输入的字符串，处理 空格　- + 及非数字字符
      step2: 遍历规范后的字符串计算结果
      step3: 根据 - + 判断是否越界并返回正确结果
 */
func myAtoiMine(str string) int {
	trimspacestr := strings.TrimSpace(str)
	strbytes := []byte(trimspacestr)
	nagflag := false
	var buffer bytes.Buffer
	if len(strbytes) == 0 {
		return 0
	} else if strbytes[0] == '-' && len(strbytes) > 1 {
		nagflag = true
		trimspacestr = strings.TrimLeft(string(strbytes[1:]), "0")
	} else if strbytes[0] == '+' && len(strbytes) > 1 {
		nagflag = false
		trimspacestr = strings.TrimLeft(string(strbytes[1:]), "0")
	}

	//log.Println("trimspacestr=", trimspacestr)
	trimspacestr = strings.TrimLeft(trimspacestr, "0")
	strbytes = []byte(trimspacestr)
	if len(strbytes) == 0 || strbytes[0] < '1' || strbytes[0] > '9' {
		return 0
	}

	//log.Println("strbytes=", strbytes)
	for _, val := range strbytes {
		if val >= '0' && val <= '9' {
			buffer.WriteByte(val)
		} else {
			break
		}
	}
	rbytes := buffer.Bytes()
	//log.Println("rbytes=", string(rbytes))
	result := 0
	for _,val := range rbytes {
		value := int(val-'0')
		result *= 10
		result += value
		if nagflag && result > 2147483648 {
			result = 2147483648
			break
		} else if !nagflag && result > 2147483647 {
			result = 2147483647
			break
		}
	}

	if nagflag {
		return int(-result)
	} else {
		return int(result)
	}
}

/**
https://leetcode-cn.com/problems/string-to-integer-atoi/solution/dan-zhi-zhen-bian-li-chuan-by-pppobear/

算法：
单指针遍历串
step1: 从字符串头开始遍历，记录有效起始位置为i，遇到 + - 则跳转到下一个字符查找末尾位置，如果遇到 空格 则继续遍历找到第一个非空格字符，其他情况返回0
step2: 从位置i开始遍历查找末尾位置j，直到遇到第一个非数字字符
step3: 从位置i遍历到j计算结果，并根据 + - 判断溢出和返回最终结果
 */
const (
	int32Max = 1 << 31 - 1
	int32Min = -1 << 31
)
func myAtoiLCA(str string) int {
	strlen := len(str)
	nav := false
	result := 0
	i, j, k := 0, 0, 0
	for i = 0; i < strlen; i++ {
		if str[i] == '+' {
			i++
			break
		} else if str[i] == '-' {
			i++
			nav = true
			break
		} else if str[i] >= '0' && str[i] <= '9' {
			break
		} else if str[i] != ' ' {
			return 0
		}
	}

	for j = i; j < strlen; j++ {
		if str[j] < '0' || str[j] > '9' {
			break
		}
	}

	for k = i; k < j; k++ {
		result *= 10
		cur := int(str[k]-'0')

		if nav {
			result -= cur
			if result < int32Min {
				return int32Min
			}
		} else {
			result += cur
			if result > int32Max {
				return int32Max
			}
		}
	}

	return result
}