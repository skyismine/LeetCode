package problems

/**
题目描述：
    判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

    示例 1:

    输入: 121
    输出: true
    示例 2:

    输入: -121
    输出: false
    解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
    示例 3:

    输入: 10
    输出: false
    解释: 从右向左读, 为 01 。因此它不是一个回文数。

    进阶:

    你能不将整数转为字符串来解决这个问题吗？

https://leetcode-cn.com/problems/palindrome-number/
 */
func IsPalindrome(x int) bool {
	switch METHOD {
	case METHOD_MINE:
		return isPalindromeMine(x)
	case METHOD_LCA:
		return isPalindromeLCA(x)
	}
	return false
}

/**
 自己想的方法：将数字每一位取出存入数组，对数组中的元素进行对折比较，即 0 和 length-1 比
 以此类推，如果有不相等的就说明不是回文
 空间和时间复杂度都为 O(log10(x))
 */
func isPalindromeMine(x int) bool {
	//负数前面有符号 -，所以不可能是回文，最后一位为0且为回文的数只可能是0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var numarrays []byte
	for x != 0 {
		numarrays = append(numarrays, byte(x%10))
		x /= 10
	}

	length := len(numarrays)
	for i := 0; i < length; i++ {
		if numarrays[i] != numarrays[length-1-i] {
			return false
		}
	}

	return true
}

/**
算法
	首先，我们应该处理一些临界情况。所有负数都不可能是回文，例如：-123 不是回文，因为 - 不等于 3。所以我们可以对所有负数返回 false。
	现在，让我们来考虑如何反转后半部分的数字。 对于数字 1221，如果执行 1221 % 10，我们将得到最后一位数字 1，要得到倒数第二位数字，我们可以先通过除以 10 把最后一位数字从 1221 中移除，1221 / 10 = 122，再求出上一步结果除以10的余数，122 % 10 = 2，就可以得到倒数第二位数字。如果我们把最后一位数字乘以10，再加上倒数第二位数字，1 * 10 + 2 = 12，就得到了我们想要的反转后的数字。 如果继续这个过程，我们将得到更多位数的反转数字。
	现在的问题是，我们如何知道反转数字的位数已经达到原始数字位数的一半？
	我们将原始数字除以 10，然后给反转后的数字乘上 10，所以，当原始数字小于反转后的数字时，就意味着我们已经处理了一半位数的数字。

复杂度分析
	时间复杂度：O(log10(n))， 对于每次迭代，我们会将输入除以10，因此时间复杂度为 O(log10(n))。
	空间复杂度：O(1)。
 */
func isPalindromeLCA(x int) bool {
	//负数前面有符号 -，所以不可能是回文，最后一位为0且为回文的数只可能是0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	//当原始数字小于反转后的数字时，就意味着我们已经处理了一半位数的数字。
	var temp int = 0
	for x > temp {
		temp = temp*10+x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	if x == temp || x == temp/10 {
		return true
	}
	return false
}

/**
 leetcode上摘取的执行速度最快的程序，算法和lca差不多，只是程序处理上有优化

复杂度分析
	时间复杂度：O(log10(n))， 对于每次迭代，我们会将输入除以10，因此时间复杂度为 O(log10(n))。
	空间复杂度：O(1)。
 */
func isPalindromeOthers(x int) bool {
	//负数前面有符号 -，所以不可能是回文，最后一位为0且为回文的数只可能是0
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	}

	//当原始数字小于反转后的数字时，就意味着我们已经处理了一半位数的数字。
	var temp int = 0
	for x > temp {
		temp = temp*10+x%10
		x /= 10
		//比较操作放在for循环里面有可能会少执行一次循环
		if x == temp || x/10 == temp {
			return true
		}
	}

	return false
}