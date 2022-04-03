package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 众数是指一组数据中出现次数量多的那个数，众数可以是多个

// 中位数是指把一组数据从小到大排列，最中间的那个数，如果这组数据的个数是奇数，那最中间那个就是中位数，如果这组数据的个数为偶数，那就把中间的两个数之和除以2，所得的结果就是中位数

// 查找整型数组中元素的众数并组成一个新的数组，求新数组的中位数

// 输入描述:

// 输入一个一维整型数组，数组大小取值范围 0<N<1000，数组中每个元素取值范围 0<E<1000

// 输出描述:

// 输出众数组成的新数组的中位数

// 示例1：

// 输入

// 10 11 21 19 21 17 21 16 21 18 15

// 输出

// 21

// 示例2：

// 输入

// 2 1 5 4 3 3 9 2 7 4 6 2 15 4 2 4
// 输出

// 3

// 示例3：

// 输入

// 5 1 5 3 5 2 5 5 7 6 7 3 7 11 7 55 7 9 98 9 17 9 15 9 9 1 39

// 输出

// 7

type numDesc struct {
	num int
	n   int
}
type numArr []numDesc

func (arr numArr) Len() int {
	return len(arr)
}

func (arr numArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr numArr) Less(i, j int) bool {
	return arr[i].n < arr[j].n

}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		input := scan.Text()
		arr := strings.Split(input, " ")
		m := make(map[string]int)
		for _, v := range arr {
			m[v]++
		}
		r := numArr{}
		for k, v := range m {
			_k, _ := strconv.Atoi(k)
			tmp := numDesc{_k, v}
			r = append(r, tmp)
		}
		sort.Sort(r)
		res := 0
		num := []int{}

		num = append(num, r[len(r)-1].num)
		for j := len(r) - 2; j >= 0; j-- {
			if r[len(r)-1].n == r[j].n {
				num = append(num, r[j].num)
			} else {
				break
			}
		}
		sort.Ints(num)
		fmt.Println("nums=", num)
		if len(num)%2 == 0 {
			res = (num[len(num)/2] + num[len(num)/2-1]) / 2
		} else {
			res = num[len(num)/2]
		}
		fmt.Println(res)

	}

}
