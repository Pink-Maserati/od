package main

import "fmt"

// 小明从糖果盒中随意抓一把糖果，每次小明会取出一半的糖果分给同学们。

// 当糖果不能平均分配时，小明可以选择从糖果盒中（假设盒中糖果足够）取出一个糖果或放回一个糖果。

// 小明最少需要多少次（取出、放回和平均分配均记一次），能将手中糖果分至只剩一颗

// 输入描述:

// 抓取的糖果数（<10000000000）：

// 15

// 输出描述:

// 最少分至一颗糖果的次数：

// 5

// 示例1：

// 输入

// 15

// 输出

// 5

// 备注:

// 解释：(1)15+1=16;(2)16/2=8;(3)8/2=4;(4)4/2=2;(5)2/2=1;

func countCal(sum, num int) int {
	//fmt.Println(sum, num)
	if sum <= 1 {
		return num
	}
	if sum%2 == 0 {
		return countCal(sum/2, num+1)
	}

	return min(countCal(sum+1, num+1), countCal(sum-1, num+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scanln(&n)
	tmp := n
	count := 0
	// for n%2 == 0 {
	// 	count++
	// 	n = n / 2
	// 	fmt.Println(n)
	// 	if n == 1 {
	// 		break
	// 	}

	// }

	flag := true
	for n != 1 {
		for n%2 == 0 {
			count++
			n = n / 2
			fmt.Println(n)
			flag = false

		}
		if flag {
			if n%2 == 1 {
				// 1..........01  -> 1...........100 下次一定会减少两位
				// 1..........11  -> 1............100 下次一定会减少两位
				if n == 3 || n/2%2 == 0 {
					count++
					n = n / 2
				} else {
					n++
				}
			}
			fmt.Println(n)
			count++

		}

	}
	fmt.Println(count)

	fmt.Println(countCal(tmp, 0))
}
