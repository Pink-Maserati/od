package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//  内存资源分配

//  有一个简易内存池，内存按照大小粒度分类，每个粒度有若干个可用内存资源，用户会进行一系列内存申请，需要按需分配内存池中的资源，返回申请结果成功失败列表。

//  分配规则如下：

//  1、分配的内存要大于等于内存申请量，存在满足需求的内存就必须分配，优先分配粒度小的，但内存不能拆分使用。

//  2、需要按申请顺序分配，先申请的先分配。

//  3、有可用内存分配则申请结果为true，没有可用内存分配则返回false。

//  注：不考虑内存释放。

//  输入描述:

//  输入为两行字符串：

//  第一行为内存池资源列表，包含内存粒度数据信息，粒度数据间用逗号分割，一个粒度信息内部用冒号分割，冒号前为内存粒度大小，冒号后为数量。

//  资源列表不大于1024，每个粒度的数量不大于4096

//  第二行为申请列表，申请的内存大小间用逗号分隔。申请列表不大于100000
//  如：
// 64:2,128:1,32:4,1:128
// 50,36,64,128,127
// 输出描述:

// 输出为内存池分配结果。

// 如：

// true,true,true,false,false

// 示例 1

// 输入

// 64:2,128:1,32:4,1:128
// 50,36,64,128,127
// 输出

// true,true,true,false,false

// 样例解释

// 内存池资源包含：64K共2个、128K共1个、32K共4个、1K共128个的内存资源；

// 针对50,36,64,128,127的内存申请序列，分配的内存依次是：64,64,128,NULL,NULL,

// 第三次申请内存时已经将128分配出去，

// 因此输出结果是：

// true,true,true,false,false

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line1 := sc.Text()
	sc.Scan()
	line2 := sc.Text()
	txt1 := strings.Split(line1, ",")
	txt2 := strings.Split(line2, ",")
	arr := []int{}
	m := make(map[int]int, len(txt1))
	for i := 0; i < len(txt1); i++ {
		tmp := strings.Split(txt1[i], ":")
		k, _ := strconv.Atoi(tmp[0])
		v, _ := strconv.Atoi(tmp[1])
		m[k] = v
		arr = append(arr, k)

	}
	arr2 := []int{}
	for i := 0; i < len(txt2); i++ {
		tmp, _ := strconv.Atoi(txt2[i])
		arr2 = append(arr2, tmp)

	}
	//fmt.Println(arr, arr2, txt1, txt2, m)
	sort.Ints(arr)
	fmt.Println(arr)
	str := strings.Builder{}

	for _, v2 := range arr2 {
		flag := false
		for _, v := range arr {
			if v2 <= v && m[v] != 0 {
				m[v]--

				flag = true
				break
			}
		}
		//	fmt.Println(m)
		if flag {
			str.WriteString("true,")
		} else {
			str.WriteString("false,")
		}
	}

	fmt.Println(strings.TrimSuffix(str.String(), ","))

}
