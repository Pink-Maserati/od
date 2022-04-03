package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 某学校举行运动会，学生们按编号(1、2、3…n)进行标识，现需要按照身高由低到高排列，对身高相同的人，按体重由轻到重排列；对于身高体重都相同的人，维持原有的编号顺序关系。请输出排列后的学生编号。

// 输入描述:

// 两个序列，每个序列由n个正整数组成（0 < n <= 100）。第一个序列中的数值代表身高，第二个序列中的数值代表体重。

// 输出描述:

// 排列结果，每个数值都是原始序列中的学生编号，编号从1开始
// 示例1：

// 输入

// 4

// 100 100 120 130

// 40 30 60 50

// 输出

// 2 1 3 4

// 说明

// 输出的第一个数字2表示此人原始编号为2，即身高为100，体重为30的这个人。由于他和编号为1的人身高一样，但体重更轻，因此要排在1前面。

// 示例2：

// 输入

// 3

// 90 110 90

// 45 60 45

// 输出

// 1 3 2

// 说明
// 1和3的身高体重都相同，需要按照原有位置关系让1排在3前面，而不是3 1 2
type student struct {
	id     int
	height int
	weight int
}
type sArr []student

func (arr sArr) Len() int {
	return len(arr)
}
func (arr sArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr sArr) Less(i, j int) bool {
	if arr[i].height == arr[j].height {
		return arr[i].weight < arr[j].weight
	} else if arr[i].height == arr[j].height && arr[i].weight == arr[j].weight {
		return arr[i].id < arr[j].id
	} else {
		return arr[i].height < arr[j].height
	}
}

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line1 := sc.Text()
	n, _ := strconv.Atoi(line1)
	sc.Scan()
	line2 := sc.Text()
	arr1 := strings.Split(line2, " ")
	height := []int{}
	for i := 0; i < n; i++ {
		tmp, _ := strconv.Atoi(arr1[i])
		height = append(height, tmp)
	}
	sc.Scan()
	line3 := sc.Text()
	arr2 := strings.Split(line3, " ")
	weight := []int{}
	for i := 0; i < n; i++ {
		tmp, _ := strconv.Atoi(arr2[i])

		weight = append(weight, tmp)
	}
	arr := sArr{}
	for i := 0; i < n; i++ {
		s := student{i + 1, height[i], weight[i]}
		arr = append(arr, s)

	}
	sort.Sort(arr)
	fmt.Print(arr[0].id)
	for i := 1; i < n; i++ {
		fmt.Print(" ", arr[i].id)
	}
	fmt.Println()

}
