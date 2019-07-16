package problems

/**
问题描述：
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

https://leetcode-cn.com/problems/search-insert-position/
 */
func SearchInsert(nums []int, target int) int {
	switch METHOD {
	case METHOD_MINE:
		return searchInsertMine(nums, target)
	case METHOD_LCA:
		return searchInsertLCA(nums, target)
	}

	return 0
}

/**
算法：二分法查找
复杂度分析：
时间复杂度：O(log(N))，计算方法：
二分法的关键思想是   假设该数组的长度是N那么二分后是N/2，再二分后是N/4……直到二分到1结束
（当然这是属于最坏的情况了，即每次找到的那个中点数都不是我们要找的），那么二分的次数就是基本语句执行的次数，
于是我们可以设次数为x，N*（1/2）^x=1；则x=logn,底数是2
空间复杂度：O(1) 不占用额外空间
 */
func searchInsertMine(nums []int, target int) int {
	min, max := 0, len(nums)
	for min < max {
		middle := (min+max) >> 1
		if nums[middle] > target {
			max = middle
		} else if nums[middle] < target {
			min = middle+1
		} else {
			return middle
		}
	}

	return min
}

/**
LCA 暂无解答
 */
func searchInsertLCA(nums []int, target int) int {
	return 0
}
