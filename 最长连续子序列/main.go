package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 最长连续子序列

// 有N个正整数组成的一个序列。给定整数sum，求长度最长的连续子序列，使他们的和等于sum，返回此子序列的长度，如果没有满足要求的序列，返回-1。

// 输入描述:

// 序列：1,2,3,4,2
// sum：6
// 输出描述:

// 序列长度：3
// 输入描述:

// 序列：1,2,3,4,2
// sum：6
// 输出描述:

// 序列长度：3
// 示例1
// 输入

// 1,2,3,4,2
// 6
// 输出

// 3
// 说明

// 解释：1,2,3和4,2两个序列均能满足要求，所以最长的连续序列为1,2,3，因此结果为3

// 示例2

// 输入

// 1,2,3,4,2
// 20
// 输出

// -1
// 说明

// 解释：没有满足要求的子序列，返回-1
// 备注:

// 输入序列仅由数字和英文逗号构成，数字之间采用英文逗号分隔；
// 序列长度：1 <= N <= 200；
// 输入序列不考虑异常情况，由题目保证输入序列满足要求。
func main() {
	var line string
	fmt.Scanln(&line)
	var sum int
	fmt.Scanln(&sum)
	strs := strings.Split(line, ",")

	maxLen := -1
	for i := 0; i < len(strs); i++ {
		tmp := 0
		tmpLen := 0
		for j := i; j < len(strs); j++ {
			if tmp > sum {
				break
			}
			num, _ := strconv.Atoi(strs[j])
			tmp += num
			tmpLen++
			//	fmt.Println("长度", tmpLen, j-i+1)
			if tmp == sum && tmpLen > maxLen {
				maxLen = tmpLen
			}
		}
	}
	fmt.Println(maxLen)

}
