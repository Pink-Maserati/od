package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 小组中每位都有一张卡片，卡片上是6位内的正整数，将卡片连起来可以组成多种数字，计算组成的最大数字。

// 输入描述:

// “,”号分割的多个正整数字符串，不需要考虑非数字异常情况，小组最多25个人

// 输出描述:

// 最大的数字字符串

// 示例1

// 输入

// 22,221

// 输出

// 22221

// 示例2

// 输入

// 4589,101,41425,9999

// 输出

// 9999458941425101

func main() {

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		arr := strings.Split(txt, ",")
		sort.Slice(arr, func(i, j int) bool {
			return arr[i]+arr[j] > arr[j]+arr[i]
		})
		if arr[0] == "0" {
			fmt.Println("0")
		} else {
			res := strings.Builder{}
			for i := 0; i < len(arr); i++ {
				res.WriteString(arr[i])
			}
			fmt.Println(res.String())

		}

	}
}
