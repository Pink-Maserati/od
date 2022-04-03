package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 公司用一个字符串来表示员工的出勤信息：

// absent：缺勤

// late：迟到

// leaveearly：早退

// present：正常上班

// 现需根据员工出勤信息，判断本次是否能获得出勤奖，能获得出勤奖的条件如下：

// 缺勤不超过一次；没有连续的迟到/早退；任意连续7次考勤，缺勤/迟到/早退不超过3次

// 输入描述:

// 用户的考勤数据字符串，记录条数 >= 1；输入字符串长度<10000；不存在非法输入

// 如：

// 2

// present

// present absent present present leaveearly present absent

// 输出描述:

// 根据考勤数据字符串，如果能得到考勤奖，输出"true"；否则输出"false"，对于输入示例的结果应为：
// true false

// 示例1：

// 输入

// 2

// present

// present present

// 输出

// true true

// 示例2：

// 输入

// 2

// present

// present absent present present leaveearly present absent

// 输出

// true false

func getAbsentTimes(arr []string) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == "absent" {
			res++
		}
	}
	return res
}

func getConditionTwo(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] == "late" || arr[i] == "leaveearly") && (arr[i+1] == "late" || arr[i+1] == "leaveearly") {
			return true
		}
	}
	return false

}

func getConditionTree(arr []string) bool {
	n := len(arr)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		if arr[i] != "present" {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	if n <= 7 && getSum(nums) >= 3 {
		return true

	} else {
		for i := 0; i < n-7; i++ {
			subNums := nums[i : i+7]
			if getSum(subNums) >= 3 {
				return true
			}
		}
		return false
	}

}

func getSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	var n int
	fmt.Scanln(&n)
	days := []string{}
	sc := bufio.NewScanner(os.Stdin)
	i := 0
	for sc.Scan() {
		i++
		days = append(days, sc.Text())
		if i == n {
			break
		}

	}
	res := []string{}

	for j := 0; j < n; j++ {
		tmp := days[j]
		arr := strings.Split(tmp, " ")
		//判断条件1：获取缺勤次数
		absentTimes := getAbsentTimes(arr)
		if absentTimes > 1 {
			res = append(res, "false")
			continue
		}
		//条件2：没有连续的迟到/早退
		condition2 := getConditionTwo(arr)
		if condition2 {
			res = append(res, "false")

		}
		//条件3：任意连续7次考勤，缺勤/迟到/早退不超过3次
		condition3 := getConditionTree(arr)
		if condition3 {
			res = append(res, "false")
		}
		res = append(res, "true")

	}

	fmt.Println(strings.Join(res, " "))
}
