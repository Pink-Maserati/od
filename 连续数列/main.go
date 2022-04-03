package main

import "fmt"

// 已知连续正整数数列{K}=K1,K2,K3...Ki的各个数相加之和为S，i=N (0<S<100000, 0<N<100000), 求此数列K。

// 输入描述:

// 输入包含两个参数，1）连续正整数数列和S，2）数列里数的个数N。

// 输出描述:

// 如果有解输出数列K，如果无解输出-1

// 示例1
// 输入
// 525 6
// 输出
// 85 86 87 88 89 90

// 示例2
// 输入
// 3 5
// 输出
// -1

func main() {
	var s int
	var n int
	fmt.Scanf("%d", &s)
	fmt.Scanf("%d", &n)
	//  s = (a1 + an)*n/2-------->  a1 = (2s/n +1 -n )/2  ,则 2s/n - n 必须为奇数

	if 2*s%n != 0 {
		fmt.Println(-1)

	} else if (2*s/n-n)%2 == 0 {
		fmt.Println(-1)
	} else {
		a1 := (2*s/n + 1 - n) / 2
		for i := 0; i < n; i++ {
			if i != n-1 {
				fmt.Print(a1 + i)
				fmt.Print(" ")
			} else {
				fmt.Print(a1 + i)
				fmt.Println()
			}

		}
	}

}
