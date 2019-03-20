package problems

import "fmt"

func RomanToInt(s,method string) int {
	switch method {
	case "mine":
		return romanToIntMine(s)
	case "mine2":
		return romanToIntMine2(s)
	case "lca":
		return romanToIntLCA(s)
	}

	return 0
}

/**
 遍历字符串对于I X C字符的处理需要结合它的下一个字符来进行
 复杂度分析：
     时间复杂度为O(n)
     空间复杂度为O(n)
 */
func romanToIntMine(s string) int {
	strbytes := []byte(s)
	strlen := len(strbytes)
	ret := 0
	for i := 0; i < strlen; i++ {
		switch strbytes[i] {
		case 'I':
			if i >= strlen-1 {
				ret += 1
			} else {
				if strbytes[i+1] == 'V' {
					ret += 4
					i++
				} else if (strbytes[i+1] == 'X') {
					ret += 9
					i++
				} else {
					ret += 1
				}
			}
		case 'V':
			ret += 5
		case 'X':
			if i >= strlen-1 {
				ret += 10
			} else {
				if strbytes[i+1] == 'L' {
					ret += 40
					i++
				} else if (strbytes[i+1] == 'C') {
					ret += 90
					i++
				} else {
					ret += 10
				}
			}
		case 'L':
			ret += 50
		case 'C':
			if i >= strlen-1 {
				ret += 100
			} else {
				if strbytes[i+1] == 'D' {
					ret += 400
					i++
				} else if (strbytes[i+1] == 'M') {
					ret += 900
					i++
				} else {
					ret += 100
				}
			}
		case 'D':
			ret += 500
		case 'M':
			ret += 1000
		}
	}
	return ret
}

/**
 将可能的值都存储在map中，利用I*100+V X*100+L等几个可能的值不相等不会导致key冲突的特性使用int作为键，遍历字符串如果遇到I X C则需要结合下一个字符
共同处理，否则直接加当前处理结果
 复杂度分析：
     时间复杂度为O(n)
     空间复杂度为O(n)
  没有 romanToIntMine 效率高，具体原因待分析
 */
func romanToIntMine2(s string) int {
	ixcmap := map[int]int{
		'I' : 1,
		'I'*100+'V' : 4,
		'I'*100+'X' : 9,
		'V' : 5,
		'X' : 10,
		'X'*100+'L' : 40,
		'X'*100+'C' : 90,
		'L' : 50,
		'C' : 100,
		'C'*100+'D' : 400,
		'C'*100+'M' : 900,
		'D' : 500,
		'M' : 1000,
	}
	fmt.Println(ixcmap)
	strbytes := []byte(s)
	strlen := len(strbytes)
	ret := 0
	for i := 0; i < strlen; i++ {
		var curr int = int(strbytes[i])
		if i < strlen-1 && (curr == 'I' || curr == 'X' || curr == 'C') {
			var next int = int(strbytes[i+1])
			fmt.Println(fmt.Sprintf("curr: %c, next: %c, sum: %d", curr, next, curr*100+next))
			if v, ok := ixcmap[curr*100+next]; ok {
				ret += v
				i++
				continue
			}
		}

		ret += ixcmap[int(curr)]
	}
	return ret
}

/**
 LCA 上没有解答
 */
func romanToIntLCA(s string) int {
	return 0
}
