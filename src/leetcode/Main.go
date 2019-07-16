package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"problems"
	"runtime/pprof"
	"time"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	method := flag.String("method", problems.METHOD_MINE, "alth method")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	problems.SetMethod(*method)

	start := time.Now()
	/*
	fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9))
	fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9))
	fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9))
	fmt.Println(problems.Reverse(1534))
	fmt.Println(problems.Reverse(1534))
	fmt.Println(problems.IsPalindrome(121))
	fmt.Println(problems.IsPalindrome(1234))
	fmt.Println(problems.IsPalindrome(10))
	fmt.Println(problems.IsPalindrome(math.MaxInt32))
	fmt.Println(problems.IsPalindrome(1221))
	fmt.Println(problems.RomanToInt("III"))
	fmt.Println(problems.RomanToInt("IV"))
	fmt.Println(problems.RomanToInt("IX"))
	fmt.Println(problems.RomanToInt("LVIII"))
	fmt.Println(problems.RomanToInt("MCMXCIV"))
	fmt.Println(problems.RomanToInt("DCXXI"))
	fmt.Println(problems.LongestCommonPrefix(nil))
	fmt.Println(problems.LongestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(problems.LongestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(problems.LongestCommonPrefix(nil))
	fmt.Println(problems.LongestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(problems.LongestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(problems.IsValid("([]"))
	fmt.Println(problems.IsValid("()"))
	fmt.Println(problems.IsValid("()[]{}"))
	fmt.Println(problems.IsValid("(]"))
	fmt.Println(problems.IsValid("([)]"))
	fmt.Println(problems.IsValid("{[]}"))
	fmt.Println(problems.IsValid("([]"))
	fmt.Println(problems.IsValid("()"))
	fmt.Println(problems.IsValid("()[]{}"))
	fmt.Println(problems.IsValid("(]"))
	fmt.Println(problems.IsValid("([)]"))
	fmt.Println(problems.IsValid("{[]}"))
	l1 := &problems.ListNode{1, &problems.ListNode{2, &problems.ListNode{4, nil}}}
	l2 := &problems.ListNode{1, &problems.ListNode{3, &problems.ListNode{4, nil}}}
	l3 := problems.MergeTwoLists(l1, l2)
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{2, nil}
	l2 = &problems.ListNode{1, nil}
	l3 = problems.MergeTwoLists(l1, l2)
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{5, nil}
	l2 = &problems.ListNode{1,  &problems.ListNode{2,  &problems.ListNode{4, nil}}}
	l3 = problems.MergeTwoLists(l1, l2)
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{-10, &problems.ListNode{-9, &problems.ListNode{-6, &problems.ListNode{-4, &problems.ListNode{1, &problems.ListNode{9, &problems.ListNode{9, nil}}}}}}}
	l2 = &problems.ListNode{-5,  &problems.ListNode{-3,  &problems.ListNode{0, &problems.ListNode{7, &problems.ListNode{8, &problems.ListNode{8, nil}}}}}}
	l3 = problems.MergeTwoLists(l1, l2)
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	arr := []int{1,1,2}
	lendup := 0
	lendup = problems.RemoveDuplicates(arr)
	lendup = problems.RemoveDuplicates(arr)
	fmt.Println(lendup, arr)
	arr = []int{0,0,1,1,1,2,2,3,3,4}
	lendup = problems.RemoveDuplicates(arr)
	lendup = problems.RemoveDuplicates(arr)
	fmt.Println(lendup, arr)
	nums := []int{3,2,2,3}
	lenele := problems.RemoveElement(nums, 3)
	fmt.Println(lenele, nums)
	nums = []int{0,1,2,2,3,0,4,2}
	lenele = problems.RemoveElement(nums, 2)
	fmt.Println(lenele, nums)
	fmt.Println(problems.StrStr("hello", "ll"))
	fmt.Println(problems.StrStr("aaaaa", "bba"))
	fmt.Println(problems.StrStr("a", "a"))
	fmt.Println(problems.StrStr("mississippi", "issip"))
	fmt.Println(problems.SearchInsert([]int{1,3,5,6}, 5))
	fmt.Println(problems.SearchInsert([]int{1,3,5,6}, 2))
	fmt.Println(problems.SearchInsert([]int{1,3,5,6}, 7))
	fmt.Println(problems.SearchInsert([]int{1,3,5,6}, 0))
	fmt.Println(problems.CountAndSay(1))
	fmt.Println(problems.CountAndSay(2))
	fmt.Println(problems.CountAndSay(3))
	fmt.Println(problems.CountAndSay(4))
	fmt.Println(problems.CountAndSay(5))
	fmt.Println(problems.CountAndSay(6))
	fmt.Println(problems.CountAndSay(7))
	fmt.Println(problems.CountAndSay(8))
	fmt.Println(problems.CountAndSay(9))
	fmt.Println(problems.CountAndSay(10))
	ll1 := &problems.ListNode{7, &problems.ListNode{8, &problems.ListNode{8, nil}}}
	ll2 := &problems.ListNode{5,  &problems.ListNode{6,  &problems.ListNode{4, nil}}}
	printLinkedList([]*problems.ListNode{problems.AddTwoNumbers(ll1, ll2)})
	fmt.Println(problems.LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(problems.LengthOfLongestSubstring("bbbbb"))
	fmt.Println(problems.LengthOfLongestSubstring("pwwkew"))
	fmt.Println(problems.FindMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(problems.FindMedianSortedArrays([]int{1,2}, []int{3,4}))
	fmt.Println(problems.FindMedianSortedArrays([]int{1,1,1,1,1,1,1,1,1,1,4,4}, []int{1,3,4,4,4,4,4,4,4,4,4}))
	fmt.Println(problems.LongestPalindrome("babad"))
	fmt.Println(problems.LongestPalindrome("bb"))
	fmt.Println(problems.LongestPalindrome("a"))
	fmt.Println(problems.LongestPalindrome("ac"))
	fmt.Println(problems.LongestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(problems.LongestPalindrome("abcba"))
	fmt.Println(problems.ZigzagConvert("LEETCODEISHIRING", 3))
	fmt.Println(problems.ZigzagConvert("LEETCODEISHIRING", 4))
	fmt.Println(problems.ZigzagConvert("A", 1))
	*/
	fmt.Println(problems.MyAtoi("42"))
	fmt.Println(problems.MyAtoi("-42"))
	fmt.Println(problems.MyAtoi("4193 with words"))
	fmt.Println(problems.MyAtoi("words and 987"))
	fmt.Println(problems.MyAtoi("-91283472332"))
	fmt.Println(problems.MyAtoi("  0000000000012345678"))
	fmt.Println(problems.MyAtoi("-000000000000001"))
	fmt.Println(problems.MyAtoi(" 2048 "))
	fmt.Println(problems.MyAtoi("+02048 "))
	fmt.Println(problems.MyAtoi("+1"))
	fmt.Println(problems.MyAtoi("1"))
	fmt.Println(problems.MyAtoi("0-1"))
	fmt.Println(problems.MyAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
	fmt.Println(problems.MyAtoi("+004500"))
	defer func() {
		fmt.Println("cost time:", time.Since(start))
	}()
	/*select {
	case <-time.After(10*time.Second):
		break
	}*/
}

func printLinkedList(la []*problems.ListNode) {
	for _, value := range la {
		linkedlist := value
		for linkedlist != nil {
			fmt.Printf("%d\t", linkedlist.Val)
			linkedlist = linkedlist.Next
		}
		fmt.Println()
	}
}
