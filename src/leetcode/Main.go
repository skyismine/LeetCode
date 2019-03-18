package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	fmt.Println(problems.IsPalindrome(10, "mine"))*/
	fmt.Println(problems.IsPalindrome(math.MaxInt32, "lca"))
	fmt.Println(problems.IsPalindrome(1221, "lca"))
	defer func() {
		fmt.Println("cost time:", time.Since(start))
	}()
}
