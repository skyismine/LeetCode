package problems

/**
题目描述:

盛最多水的容器

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49

链接：https://leetcode-cn.com/problems/container-with-most-water
 */
func ContainerWithMostWater(height []int) int {
	switch METHOD {
	case METHOD_MINE:
		return containerWithMostWaterMine(height)
	case METHOD_LCA:
		return containerWithMostWaterLCA(height)
	}

	return 0
}

/**
思路: 两次遍历数组挨个计算container并保存大的
复杂度分析：
时间复杂度: O(n^2)
空间复杂度: O(1)
 */
func containerWithMostWaterMine(height []int) int {
	arrlen := len(height)
	container := 0
	for i := 0; i < arrlen; i++ {
		for j := i; j < arrlen; j++ {
			minij := height[i]
			if minij > height[j] {
				minij = height[j]
			}
			tmp := (j-i) * minij
			if tmp > container {
				container = tmp
			}
		}
	}
	return container
}

/**
思路：使用双指针法，初始时两个指针分别指向第一个和最后一个元素，以后每次都移动指向较小元素的那个指针
复杂度分析：
时间复杂度：O(n)
空间复杂度：O(1)
 */
func containerWithMostWaterLCA(height []int) int {
	container := 0
	arrlen := len(height)
	indexf := 0
	indexs := arrlen-1
	for indexf != indexs {
		first := height[indexf]
		second := height[indexs]
		tmp := 0
		if (first <= second) {
			tmp = first * (indexs - indexf)
			indexf += 1
		} else {
			tmp = second * (indexs - indexf)
			indexs -= 1
		}
		if tmp > container {
			container = tmp
		}
	}
	return container
}