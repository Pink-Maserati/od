package main

import (
	"fmt"
	"sort"
	"strings"
)

//  单词接龙

// 题目描述：

// 单词接龙的规则是：

// 可用于接龙的单词首字母必须要前一个单词的尾字母相同；
// 当存在多个首字母相同的单词时，取长度最长的单词，如果长度也相等，则取字典序最小的单词；
// 已经参与接龙的单词不能重复使用。
// 现给定一组全部由小写字母组成单词数组，并指定其中的一个单词作为起始单词，进行单词接龙，
// 请输出最长的单词串，单词串是单词拼接而成，中间没有空格。
// 输入描述:

// 输入的第一行为一个非负整数，表示起始单词在数组中的索引K，0 <= K < N ；
// 输入的第二行为一个非负整数，表示单词的个数N；
// 接下来的N行，分别表示单词数组中的单词。
// 输出描述:

// 输出一个字符串，表示最终拼接的单词串。
// 示例 1：

// 输入

// 0
// 6
// word
// dd
// da
// dc
// dword
// d
// 输出

// worddwordda
// 说明

// 先确定起始单词word，再接以d开头的且长度最长的单词dword，
// 剩余以d开头且长度最长的有dd、da、dc，则取字典序最小的da，
// 所以最后输出worddwordda。
// 示例 2：

// 输入
// 4
// 6
// word
// dd
// da
// dc
// dword
// d
// 输出

// dwordda
// 说明

// 先确定起始单词dword，剩余以d开头且长度最长的有dd、da、dc，则取字典序最小的da，

// 所以最后输出dwordda。

// 备注:

// 单词个数N的取值范围为[1, 20]；

// 单个单词的长度的取值范围为[1, 30]；

func main() {
	var k, n int
	fmt.Scanln(&k)
	fmt.Scanln(&n)
	str := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&str[i])
	}

	if k < 0 || k >= n {
		fmt.Println("input error!")
		return
	}
	s := strings.Builder{}
	s.WriteString(str[k])
	//从数组中移除掉这个元素
	str = append(str[:k], str[k+1:]...)

	fmt.Println("排序前", str)
	//排序
	sort.Slice(str, func(i, j int) bool {
		n1 := len(str[i])
		n2 := len(str[j])
		if n1 != n2 {
			return n1 > n2
		} else {
			return str[i] < str[j]
		}
	})
	fmt.Println("排序后", str)

	for {
		//	fmt.Println("before", s.String())
		ss := s.String()
		tail := ss[len(ss)-1]
		flag := true
		for i := 0; i < len(str); i++ {
			head := str[i][0]
			if tail == head {
				s.WriteString(str[i])
				str = append(str[:i], str[i+1:]...)
				flag = false
				break
			}
		}

		//	fmt.Println("after", s.String())
		if flag {
			break
		}

	}
	fmt.Println(s.String())

}
