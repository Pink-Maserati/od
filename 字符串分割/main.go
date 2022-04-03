package main

import (
	"fmt"
	"strings"
)

// 给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。

// 输入描述:

// 输入为两行，第一行为参数K，第二行为字符串S。

// 输出描述:

// 输出转换后的字符串。

// 示例1

// 输入

// 3

// 12abc-abCABc-4aB@

// 输出

// 12abc-abc-ABC-4aB-@

// 说明

// 子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
// 示例2

// 输入

// 12

// 12abc-abCABc-4aB@

// 输出

// 12abc-abCABc4aB@

// 说明

// 子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@
func convertString(str string) string {
	little := 0
	large := 0

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			little++
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			large++
		}

	}
	var res string
	if little > large {
		res = strings.ToLower(str)

	} else if little < large {
		res = strings.ToUpper(str)
	} else {
		res = str
	}
	return res

}
func main() {
	var k int
	fmt.Scanln(&k)
	var str string
	fmt.Scanln(&str)
	arr := strings.Split(str, "-")
	n := len(arr)
	if n == 1 {
		fmt.Println(arr[0])
	} else {
		res := []string{}
		res = append(res, arr[0])
		s := strings.Join(arr[1:], "")
		left := 0
		right := left + k
		for right < len(s) {
			tmp := s[left:right]
			//需不要转化
			txt := convertString(tmp)
			res = append(res, txt)
			left += k
			right = left + k
		}
		if left < len(s) {
			res = append(res, s[left:])
		}
		fmt.Println(strings.Join(res, "-"))

	}

}
