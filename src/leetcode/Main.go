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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	start := time.Now()
	/*fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9, "baoli"))
	fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9, "hash2"))
	fmt.Println(problems.TwoSum([]int{11, 15, 18, 20, -3, -5, 90, 2, 7}, 9, "hash1"))
	fmt.Println(problems.Reverse(1534, "lca"))
	fmt.Println(problems.Reverse(1534, "mine"))
	fmt.Println(problems.IsPalindrome(121, "mine"))
	fmt.Println(problems.IsPalindrome(1234, "mine"))
	fmt.Println(problems.IsPalindrome(10, "mine"))
	fmt.Println(problems.IsPalindrome(math.MaxInt32, "lca"))
	fmt.Println(problems.IsPalindrome(1221, "lca"))
	fmt.Println(problems.RomanToInt("III", "mine2"))
	fmt.Println(problems.RomanToInt("IV", "mine2"))
	fmt.Println(problems.RomanToInt("IX", "mine2"))
	fmt.Println(problems.RomanToInt("LVIII", "mine2"))
	fmt.Println(problems.RomanToInt("MCMXCIV", "mine2"))
	fmt.Println(problems.RomanToInt("DCXXI", "mine2"))
	fmt.Println(problems.LongestCommonPrefix(nil, "mine"))
	fmt.Println(problems.LongestCommonPrefix([]string{"flower","flow","flight"}, "mine"))
	fmt.Println(problems.LongestCommonPrefix([]string{"dog","racecar","car"}, "mine"))
	fmt.Println(problems.LongestCommonPrefix(nil, "lca"))
	fmt.Println(problems.LongestCommonPrefix([]string{"flower","flow","flight"}, "lca"))
	fmt.Println(problems.LongestCommonPrefix([]string{"dog","racecar","car"}, "lca"))
	fmt.Println(problems.IsValid("([]", "mine"))
	fmt.Println(problems.IsValid("()", "mine"))
	fmt.Println(problems.IsValid("()[]{}", "mine"))
	fmt.Println(problems.IsValid("(]", "mine"))
	fmt.Println(problems.IsValid("([)]", "mine"))
	fmt.Println(problems.IsValid("{[]}", "mine"))
	fmt.Println(problems.IsValid("([]", "lca"))
	fmt.Println(problems.IsValid("()", "lca"))
	fmt.Println(problems.IsValid("()[]{}", "lca"))
	fmt.Println(problems.IsValid("(]", "lca"))
	fmt.Println(problems.IsValid("([)]", "lca"))
	fmt.Println(problems.IsValid("{[]}", "lca"))
	l1 := &problems.ListNode{1, &problems.ListNode{2, &problems.ListNode{4, nil}}}
	l2 := &problems.ListNode{1, &problems.ListNode{3, &problems.ListNode{4, nil}}}
	l3 := problems.MergeTwoLists(l1, l2, "mine")
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{2, nil}
	l2 = &problems.ListNode{1, nil}
	l3 = problems.MergeTwoLists(l1, l2, "mine")
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{5, nil}
	l2 = &problems.ListNode{1,  &problems.ListNode{2,  &problems.ListNode{4, nil}}}
	l3 = problems.MergeTwoLists(l1, l2, "mine")
	printLinkedList([]*problems.ListNode{l1, l2, l3})
	l1 = &problems.ListNode{-10, &problems.ListNode{-9, &problems.ListNode{-6, &problems.ListNode{-4, &problems.ListNode{1, &problems.ListNode{9, &problems.ListNode{9, nil}}}}}}}
	l2 = &problems.ListNode{-5,  &problems.ListNode{-3,  &problems.ListNode{0, &problems.ListNode{7, &problems.ListNode{8, &problems.ListNode{8, nil}}}}}}
	l3 = problems.MergeTwoLists(l1, l2, "mine")
	printLinkedList([]*problems.ListNode{l1, l2, l3})*/
	arr := []int{1,1,2}
	len := 0
	len = problems.RemoveDuplicates(arr, "mine")
	len = problems.RemoveDuplicates(arr, "lca")
	fmt.Println(len, arr)
	arr = []int{0,0,1,1,1,2,2,3,3,4}
	len = problems.RemoveDuplicates(arr, "mine")
	len = problems.RemoveDuplicates(arr, "lca")
	fmt.Println(len, arr)
	defer func() {
		fmt.Println("cost time:", time.Since(start))
	}()
	/*select {
	case <-time.After(10*time.Second):
		break
	}*/
}

func printLinkedList(la []*problems.ListNode) {
	for _,value := range la {
		linkedlist := value
		for linkedlist != nil  {
			fmt.Print(linkedlist.Val, "\t")
			linkedlist = linkedlist.Next
		}
		fmt.Println()
	}
}
