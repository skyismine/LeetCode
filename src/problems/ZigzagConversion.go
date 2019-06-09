package problems

import (
	"bytes"
)

/**
题目描述：
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
*/
func ZigzagConvert(s string, numRows int) string {
	return zigzagConvertMine(s, numRows)
}

/**
通过分析我们不难发现下标规律：
行数为四的时候：

第一行：0 - 6 - 12 - 18 ...     //步长6 6 6
第二行：1 - 5 - 7 - 11 - 13 ... //步长4 2 4 2
第三行：2 - 4 - 8 - 10 - 14 ... //步长2 4 2 4
第四行：3 - 9 - 15 - 21 ... //步长6 6 6



规律显而易见，当然读者可以自己多试试其他的例子来验证。其实根据"Z"字的特点我们也能直接分析出这样的规律。接下来就是将规律函数化。

步长：上述例子中 无论是6 还是4，2交替还是2，4交替，都是和6相关的。那么6又是怎么和层数关联呢？根据图形"Z"来分析，每加1层就会增加2步长，那么步长即为总层数*2-2，后文称之为绝对步长。
根据步长我们能很快计算第一层的所有下标。那第二层 第n层呢。第二层其实就是步长-2，2交替，第三层是步长步长-4,4交替。
提炼公式：步长- 2*(层数-1),2*(层数-1)交替。
当然，最后一层和第一层步长一直都是绝对步长

作者：sxqiong
链接：https://www.jianshu.com/p/bab29f7c79f9
来源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
*/
func zigzagConvertMine(s string, numRows int) string {
	absSteps := numRows*2 - 2
	if absSteps <= 0 {
		return s
	}
	strbytes := []byte(s)
	var bytesbuf bytes.Buffer
	for i := 0; i < numRows; i++ {
		index := i
		isreplace := false
		fr := absSteps - 2*i
		sr := 2 * i
		if fr != 0 {
			isreplace = false
		} else {
			isreplace = true
		}
		for index < len(s) {
			bytesbuf.WriteByte(strbytes[index])
			if isreplace {
				index += 2 * i
			} else {
				index += absSteps - 2*i
			}
			if fr != 0 && sr != 0 {
				isreplace = !isreplace
			}
		}
	}
	return bytesbuf.String()
}
