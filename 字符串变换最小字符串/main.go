package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 字符串变换最小字符串

// 给定一个字符串s，最多只能进行一次变换，返回变换后能得到的最小字符串（按照字典序进行比较）。

// 变换规则：交换字符串中任意两个不同位置的字符。

// 输入描述:

// 一串小写字母组成的字符串s

// 输出描述:

// 按照要求进行变换得到的最小字符串

// 示例1

// 输入

// abcdef
// 输出

// abcdef
// 说明

// abcdef已经是最小字符串，不需要交换

// 示例2

// 输入

// bcdefa
// 输出

// acdefb
// 说明

// a和b进行位置交换，可以得到最小字符串

// 备注:

// s是都是小写字符组成
// 1<=s.length<=1000

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	arr := []byte(input)

	str := []string{}
	str = append(str, input)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[i], arr[j] = arr[j], arr[i]
			tmp := strings.Builder{}
			tmp.WriteString(string(arr))
			arr[i], arr[j] = arr[j], arr[i]
			str = append(str, tmp.String())
			//	fmt.Println(str)

		}

	}
	sort.Strings(str)
	fmt.Println(str[0])
}
