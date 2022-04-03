package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 一辆运送快递的货车，运送的快递均放在大小不等的长方体快递盒中，为了能够装载更多的快递，同时不能让货车超载，需要计算最多能装多少个快递。

// 注：快递的体积不受限制，快递数最多1000个，货车载重最大50000。

// 输入描述:

// 第一行输入每个快递的重量，用英文逗号分隔，如：5,10,2,11

// 第二行输入货车的载重量，如：20

// 不需要考虑异常输入。

// 输出描述:

// 输出最多能装多少个快递，如：3

// 示例1：

// 输入

// 5,10,2,11

// 20

// 输出

// 3

// 说明

// 货车的载重量为20，最多只能放三个快递5、10、2，因此输出3

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line1 := sc.Text()
	arr := strings.Split(line1, ",")
	nums := []int{}
	for i := 0; i < len(arr); i++ {
		tmp, _ := strconv.Atoi(arr[i])
		nums = append(nums, tmp)
	}
	sort.Ints(nums)
	sc.Scan()
	line2 := sc.Text()
	load, _ := strconv.Atoi(line2)
	sum := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		sum += nums[i]
		if sum > load {
			break
		} else {
			count++
		}
	}
	fmt.Println(count)
}
