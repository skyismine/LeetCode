package problems

/**
 题目描述:
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

	你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]

	https://leetcode-cn.com/problems/two-sum/
 */
func TwoSum(nums []int, target int) []int {
	switch METHOD {
	case METHOD_MINE:
		return twoSumBaoli(nums, target)
	case METHOD_MINE2:
		return twoSumHash2(nums, target)
	case METHOD_LCA:
		return twoSumHash1(nums, target)
	}

	return []int{-1, -1}
}

/**
 暴力法：遍历每个元素进行计算比较
        时间复杂度：O(n^2) 对于每个元素，我们试图通过遍历数组的其余部分来寻找它所对应的目标元素，这将耗费O(n)的时间，因此总的时间复杂度为O(n^2)
        空间复杂度：O(1) 不需要额外空间，所以空间复杂度为 O(1)
 */
func twoSumBaoli(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

/**
 两遍哈希表: 第一遍构造哈希表，第二遍进行查找
            时间复杂度: O(n) 我们把包含有 n 个元素的列表遍历两次，由于哈希表将查找时间缩短到 O(1)，所以时间复杂度为 O(n)
            空间复杂度: O(n) 所需的额外空间取决于哈希表中存储的元素数量，该表中存储了 n 个元素

 这里有个bug:如果数组中有值相同的元素则构建哈希表的时候会出现冲突，这种情况没有处理
 */
func twoSumHash2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for index, value := range nums {
		numsMap[value] = index
	}

	for index, value := range nums {
		idx, ok := numsMap[target-value]
		if ok && index != idx {
			return []int{index, idx}
		}
	}
	return []int{-1, -1}
}

/**
 一遍哈希表：在进行迭代并将元素插入到哈希表中的同时检查是否已经存在目标元素
            时间复杂度: O(n) 我们只遍历数组一次，且查找元素花费的时间为 O(1)
            空间复杂度: O(n) 所需的额外空间取决于哈希表中存储的元素数量，该表中存储了 n 个元素

这里有个bug:如果数组中有值相同的元素则构建哈希表的时候会出现冲突，这种情况没有处理
 */
func twoSumHash1(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for index, value := range nums {
		idx, ok := numsMap[target-value]
		if ok && index != idx {
			return []int{idx, index}
		}
		numsMap[value] = index
	}
	return []int{-1, -1}
}