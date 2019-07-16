package problems

import (
	"math"
	"strconv"
)

/**
 题目描述：
  给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

	示例 1:

	输入: 123
	输出: 321
	 示例 2:

	输入: -123
	输出: -321
	示例 3:

	输入: 120
	输出: 21
	注意:

	假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

  https://leetcode-cn.com/problems/reverse-integer/
 */
func Reverse(x int) int {
	switch METHOD {
	case METHOD_MINE:
		return reverseMine(x)
	case METHOD_LCA:
		return reverseLCA(x)
	}
	return 0
}

/**
 自己想到的方法，将 int 转换为字符串，然后将字符数组翻转再转回 int,需要使用库函数
 时间复杂度和空间复杂度均为 O(log10(x)), x 中大约有 log10(x) 位数字
 */
func reverseMine(x int) int {
	strux := strconv.Itoa(x)
	bytesuxorig := []byte(strux)
	bytesux := bytesuxorig
	if x < 0 {
		bytesux = bytesuxorig[1:]
	}
	harflen := len(bytesux)/2
	for i := 0; i < harflen; i++ {
		j := len(bytesux)-i-1
		bytesux[i], bytesux[j] = bytesux[j], bytesux[i]
	}
	ret, err := strconv.Atoi(string(bytesuxorig))

	if err != nil || ret > 1<<31 - 1 || ret < -(1<<31){
		return 0
	}

	return ret
}

/**
 方法：弹出和推入数字 & 溢出前进行检查
思路

我们可以一次构建反转整数的一位数字。在这样做的时候，我们可以预先检查向原整数附加另一位数字是否会导致溢出。

算法

反转整数的方法可以与反转字符串进行类比。

我们想重复“弹出” x 的最后一位数字，并将它“推入”到 rev 的后面。最后，rev 将与 x 相反。

要在没有辅助堆栈 / 数组的帮助下 “弹出” 和 “推入” 数字，我们可以使用数学方法。

//pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop;
rev = temp;
但是，这种方法很危险，因为当 temp=rev⋅10+pop 时会导致溢出。

幸运的是，事先检查这个语句是否会导致溢出很容易。

为了便于解释，我们假设 rev 是正数。

如果 temp=rev⋅10+pop 导致溢出，那么一定有 rev ≥ 10*INTMAX
​
如果 rev>10*INTMAX，那么 temp=rev⋅10+pop 一定会溢出。
如果 rev==10*INTMAX 那么只要 pop>7，temp=rev⋅10+pop 就会溢出。
当 rev 为负时可以应用类似的逻辑。

复杂度分析

时间复杂度：O(log(x))，x 中大约有 log10(x) 位数字。
空间复杂度：O(1)。
 */
func reverseLCA(x int) int {
	var ret int = 0
	for x != 0 {
		mx := x%10
		x /= 10
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && mx > 7) {
			return 0
		}
		if ret < math.MinInt32/10 || (ret == math.MinInt32/10 && mx < -8) {
			return 0
		}
		ret = ret*10+mx
	}
	return ret
}