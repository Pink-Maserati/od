package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 给航天器一侧加装长方形或正方形的太阳能板（图中的红色斜线区域），需要先安装两个支柱（图中的黑色竖条），再在支柱的中间部分固定太阳能板。但航天器不同位置的支柱长度不同，太阳能板的安装面积受限于最短一侧的那根支柱长度。如图：

// 现提供一组整形数组的支柱高度数据，假设每根支柱间距离相等为1个单位长度，计算如何选择两根支柱可以使太阳能板的面积最大。

// 输入描述:

// 10,9,8,7,6,5,4,3,2,1

// 注：支柱至少有2根，最多10000根，能支持的高度范围1~10^9的整数。柱子的高度是无序的，例子中递减只是巧合。

// 输出描述:

// 可以支持的最大太阳能板面积：（10米高支柱和5米高支柱之间）

// 25

// 示例1

// 输入

// 10,9,8,7,6,5,4,3,2,1

// 输出

// 25

// 备注:

// 10米高支柱和5米高支柱之间宽度为5，高度取小的支柱高也是5，面积为25。任取其他两根支柱所能获得的面积都小于25。所以最大的太阳能板面积为25。

func maxArea(height []int) int {
	n := len(height)
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return min(height[0], height[1])
	}
	l := 0
	r := n - 1
	res := 0

	for l < r {
		tmp := min(height[l], height[r]) * (r - l)
		res = max(tmp, res)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	arr := strings.Split(input, ",")
	height := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {

		v, _ := strconv.Atoi(arr[i])
		height[i] = v
	}

	fmt.Println(maxArea(height))

}
