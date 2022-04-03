package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 给定两个整数数组array1、array2，数组元素按升序排列。假设从array1、array2中分别取出一个元素可构成一对元素，现在需要取出k对元素，并对取出的所有元素求和，计算和的最小值

// 注意：两对元素如果对应于array1、array2中的两个下标均相同，则视为同一对元素。

// 输入描述:

// 输入两行数组array1、array2，每行首个数字为数组大小size(0 < size <= 100);

// 0 < array1[i] <= 1000

// 0 < array2[i] <= 1000

// 接下来一行为正整数k

// 0 < k <= array1.size() * array2.size()

// 输出描述:

// 满足要求的最小和
// 示例1

// 输入

// 3 1 1 2

// 3 1 2 3

// 2

// 输出

// 4

// 说明

// 用例中，需要取2对元素

// 取第一个数组第0个元素与第二个数组第0个元素组成1对元素[1,1];

// 取第一个数组第1个元素与第二个数组第0个元素组成1对元素[1,1];

// 求和为1+1+1+1=4，为满足要求的最小和

//解题思路：将元素对的值，存入一个集合，对集合进行排序，然后取前k个的和

func getArray(line string) []int {
	str := strings.Split(line, " ")
	//fmt.Println(str)
	arr := make([]int, len(str)-1)
	for i := 1; i < len(str); i++ {
		tmp, _ := strconv.Atoi(str[i])
		//	fmt.Println("tmp", tmp)
		arr[i-1] = tmp
	}

	return arr

}

func minSum(arr1 []int, arr2 []int, k int) int {
	sum := []int{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			sum = append(sum, (arr1[i] + arr2[j]))

		}
	}
	sort.Ints(sum)
	min := 0
	for i := 0; i < k; i++ {
		min += sum[i]
	}
	return min

}
func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line1 := sc.Text()
	sc.Scan()
	line2 := sc.Text()
	sc.Scan()
	line3 := sc.Text()

	arr1 := getArray(line1)

	arr2 := getArray(line2)
	k, _ := strconv.Atoi(string(line3))
	fmt.Println(arr1, arr2, k)
	fmt.Println(minSum(arr1, arr2, k))

}
