package problems


/**
题目描述：

给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 */
func RemoveDuplicates(nums []int) int {
	switch METHOD {
	case METHOD_MINE:
		return removeDuplicatesMine(nums)
	case METHOD_LCA:
		return removeDuplicatesLCA(nums)
	}

	return 0
}

/**
算法：
 将数组中重复的数字移到数组最后，每遇到一个重复数字遍历的长度-1，当前元素和前一个元素不相同时下标+1继续下一项

复杂度分析：

时间复杂度：O(n^2)，n为数组长度，需要遍历n次，每次移动数组元素又需要遍历n次
空间复杂度：O(1)
 */
func removeDuplicatesMine(nums []int) int {
	lastindex := len(nums)
	for i := 1; i < lastindex; {
		if nums[i] == nums[i-1] {
			tmp := nums[i]
			lastindex--
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, tmp)
		} else {
			i++
		}
	}
	return lastindex
}

/**
算法：双指针法
数组完成排序后我们放置两个指针i和j,其中i是慢指针，而j是快指针。只要nums[i] == nums[j]，我们就
增加j以跳过重复项
当我们遇到两个不相等时则需要把nums[j]的值复制到nums[i+1]，接着递增i，直到j遍历完数组
在赋值之前需要判断i != j，避免将自己赋值给自己

复杂度分析：
时间复杂度：O(n), 假设数组的长度是n，那么i和j分别最多遍历n步
空间复杂度：O(1)
 */
func removeDuplicatesLCA(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			if i != j {
				nums[i] = nums[j]
			}
		}
	}
	return i+1
}
