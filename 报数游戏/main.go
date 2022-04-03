package main

import "fmt"

// 100个人围成一圈，每个人有一个编码，编号从1开始到100。他们从1开始依次报数，报到为M的人自动退出圈圈，然后下一个人接着从1开始报数，直到剩余的人数小于M。请问最后剩余的人在原先的编号为多少？

// 输入描述:

// 输入一个整数参数M

// 输出描述:

// 如果输入参数M小于等于1或者大于等于100，输出“ERROR!”；否则按照原先的编号从小到大的顺序，以英文逗号分割输出编号字符串

// 示例1：

// 输入

// 3

// 输出

// 58,91

// 说明

// 输入M为3，最后剩下两个人

// 示例2：

// 输入

// 4

// 输出

// 34,45,97

// 说明

// 输入M为4，最后剩下三个人

//约瑟夫问题，单循环链表（取余完成）

func Josef(m int) {
	if m <= 1 || m > 100 {
		fmt.Println("ERROR")
		return
	}
	arr := []int{}
	for i := 1; i <= 100; i++ {
		arr = append(arr, i)
	}
	index := 0
	for len(arr) >= m {
		index = (index + m - 1) % len(arr)
		arr = append(arr[:index], arr[index+1:]...)
	}
	fmt.Print(arr[0])
	for i := 1; i < len(arr); i++ {
		fmt.Print(",", arr[i])

	}
	fmt.Println()
}

func f2(m int) {
	if m <= 1 || m > 100 {
		fmt.Println("ERROR")
		return
	}
	max := 100
	marks := make([]bool, max)
	for i := 0; i < max; i++ {
		marks[i] = true
	}

	size := max
	curNum := -1
	count := 0
	for size >= m {
		curNum++
		if curNum >= max {
			curNum = curNum - max
		}

		if marks[curNum] {
			count++
			if count == m {
				marks[curNum] = false
				count = 0
				size--
			}

		}

	}

	res := []int{}
	for i := 0; i < max; i++ {
		if marks[i] {
			res = append(res, i+1)
		}

	}
	fmt.Print(res[0])
	for i := 1; i < len(res); i++ {

		fmt.Print(",", res[i])
	}
	fmt.Println()

}

func main() {
	var m int
	fmt.Scanln(&m)
	Josef(m)
	f2(m)

}
