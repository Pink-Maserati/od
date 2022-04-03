package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 分班
// 幼儿园两个班的小朋友在排队时混在了一起，每位小朋友都知道自己是否与前面一位小朋友是否同班，请你帮忙把同班的小朋友找出来。

// 小朋友的编号为整数，与前一位小朋友同班用Y表示，不同班用N表示。

// 输入描述:

// 输入为空格分开的小朋友编号和是否同班标志。

// 比如：6/N 2/Y 3/N 4/Y，表示共4位小朋友，2和6同班，3和2不同班，4和3同班。

// 其中，小朋友总数不超过999，每个小朋友编号大于0，小于等于999。

// 不考虑输入格式错误问题。

// 输出描述:

// 输出为两行，每一行记录一个班小朋友的编号，编号用空格分开。 且：

// 1.编号需要按照大小升序排列，分班记录中第一个编号小的排在第一行。
// 2.若只有一个班的小朋友，第二行为空行。
// 3.若输入不符合要求，则直接输出字符串ERROR。
// 示例1：
// 输入
// 1/N 2/Y 3/N 4/Y
// 输出
// 1 2
// 3 4
// 说明
// 2的同班标记为Y，因此和1同班。
// 3的同班标记为N，因此和1、2不同班。
// 4的同班标记为Y，因此和3同班。
// 所以1、2同班，3、4同班，输出为
// 1 2
// 3 4

func formatArray(nums []int) {
	fmt.Print(nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Print(" ", nums[1])

	}
	fmt.Println()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()
		txt := strings.Split(input, " ")
		arr1 := []int{}
		arr2 := []int{}
		flag := true
		for i := 0; i < len(txt); i++ {
			tmp := strings.Split(txt[i], "/")
			if len(tmp) == 1 {
				fmt.Println("ERROR")
				return
			}
			id := tmp[0]
			same := tmp[1]
			num, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("ERROR")
				return
			}
			if len(same) == 0 {
				fmt.Println("ERROR")
				return
			}

			if i == 0 {
				arr1 = append(arr1, num)
				continue
			}

			if same == "N" {
				flag = !flag
			}
			if flag {
				arr1 = append(arr1, num)
			} else {
				arr2 = append(arr2, num)
			}
			//fmt.Println(arr1, arr2)

		}
		//	fmt.Println(arr1, arr2)
		sort.Ints(arr1)
		sort.Ints(arr2)
		if len(arr1) < 1 {
			fmt.Println("ERROR")
			return

		}
		if len(arr2) > 0 {
			if arr1[0] < arr2[0] {
				formatArray(arr1)
				formatArray(arr2)

			} else {
				formatArray(arr2)
				formatArray(arr1)

			}
		} else {
			formatArray(arr1)
		}

	}

}
