package problems

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	switch METHOD {
	case METHOD_MINE:
		return findMedianSortedArraysMine(nums1, nums2)
	case METHOD_LCA:
		return findMedianSortedArraysLCA(nums1, nums2)
	}

	return 0
}

/**
思路：先合并两个数组，然后计算中位数
 */
func findMedianSortedArraysMine(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	var merged []int
	n1len, n2len := len(nums1), len(nums2)
	for i < n1len && j < n2len {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			merged = append(merged, nums2[j])
			j++
		} else {
			merged = append(merged, nums1[i], nums2[j])
			i++
			j++
		}
	}
	if i < n1len {
		merged = append(merged, nums1[i:]...)
	}
	if j < n2len {
		merged = append(merged, nums2[j:]...)
	}
	mlen := len(merged)
	if mlen%2 == 0 {
		return float64(merged[mlen/2-1]+merged[mlen/2])/2
	} else {
		return float64(merged[mlen/2])
	}
}

/**
方法：递归法
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/

复杂度分析：
时间复杂度：O(log(min(m,n)))，
首先，查找的区间是 [0,m]。 而该区间的长度在每次循环之后都会减少为原来的一半。 所以，我们只需要执行 log(m) 次循环。
由于我们在每次循环中进行常量次数的操作，所以时间复杂度为 O(log(m))。 由于 m≤n，所以时间复杂度是 O(log(min(m,n)))。

空间复杂度：O(1)， 我们只需要恒定的内存来存储 9 个局部变量， 所以空间复杂度为 O(1)
 */
func findMedianSortedArraysLCA(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	harf := (m+n+1)/2
	//确保nums1长度m小于nums2长度n
	if m > n {
		m,n = n,m
		nums1,nums2 = nums2,nums1
	}
	imin,imax := 0,m
	for imin <= imax {
		i := (imin+imax)/2
		j := harf-i
		if (i < imax && nums2[j-1] > nums1[i]) {
			//i is too small
			imin = i+1
		} else if (i > imin && nums1[i-1] > nums2[j]) {
			//i is too big
			imax = i-1
		} else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				return maxLeft
			}
			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			return (maxLeft+minRight)/2
		}
	}
	return 0.0
}