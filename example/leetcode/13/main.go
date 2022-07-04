package main

import "log"

func main() {
	log.Println("result", romanToInt("III"))
	log.Println("result", romanToInt("IV"))
	log.Println("result", romanToInt("IX"))
	log.Println("result", romanToInt("LVIII"))
	log.Println("result", romanToInt("MCMXCIV"))
}

//13. 罗马数字转整数
//I， V， X， L，C，D 和 M
func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var nums []int
	for _, v := range s {
		if m, ok := roman[string(v)]; ok {
			nums = append(nums, m)
		}
	}
	
	result := 0
	for i, num := range nums {
		if i == 0 {
			result += num
		} else {
			if num > nums[i - 1] {
				result = result + num - nums[i - 1] * 2
			} else {
				result += num
			}
		}
	}

	return result
}

//解题
//链接：https://leetcode-cn.com/problems/roman-to-integer/solution/luo-ma-shu-zi-zhuan-zheng-shu-by-leetcod-w55p/
//来源：力扣（LeetCode）

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt2(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

