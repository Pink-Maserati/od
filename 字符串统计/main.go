package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 给定两个字符集合，一个为全量字符集，一个为已占用字符集。已占用的字符集中的字符不能再使用，要求输出剩余可用字符集。

// 输入描述:

// 1、输入为一个字符串，一定包含@符号。@前的为全量字符集，@后的字为已占用字符集。

// 2、已占用字符集中的字符一定是全量字符集中的字符。字符集中的字符跟字符之间使用英文逗号分隔。

// 3、每个字符都表示为字符加数字的形式，用英文冒号分隔，比如a:1，表示1个a字符。

// 4、字符只考虑英文字母，区分大小写，数字只考虑正整形，数量不超过100。

// 5、如果一个字符都没被占用，@标识仍然存在，例如a:3,b:5,c:2@

// 输出描述:

// 输出可用字符集，不同的输出字符集之间回车换行。

// 注意，输出的字符顺序要跟输入一致。不能输出b:3,a:2,c:2

// 如果某个字符已全被占用，不需要再输出。
// 示例1

// 输入

// a:3,b:5,c:2@a:1,b:2

// 输出

// a:2,b:3,c:2

// 说明

// 全量字符集为3个a，5个b，2个c。

// 已占用字符集为1个a，2个b。

// 由于已占用字符不能再使用，因此，剩余可用字符为2个a，3个b，2个c。

// 因此输出a:2,b:3,c:2

func main() {

	// s := "a:3,b:5,c:2@"
	// t := strings.Split(s, "@")
	// fmt.Println(len(t), t, len(t[1]))
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()
		txts := strings.Split(input, "@")
		all := txts[0]
		fmt.Println(txts, len(txts))
		if len(txts) == 1 {
			fmt.Print(all)
			return

		} else {
			unUsed := txts[1]
			if len(unUsed) == 0 {
				fmt.Print(all)
			} else {
				unUsedArr := strings.Split(unUsed, ",")
				um := make(map[string]int)
				for _, v := range unUsedArr {
					keys := strings.Split(v, ":")
					vtmp, _ := strconv.Atoi(keys[1])
					um[keys[0]] = vtmp

				}

				allArr := strings.Split(all, ",")

				str := []string{}

				for _, v := range allArr {
					keys := strings.Split(v, ":")
					vtmp, _ := strconv.Atoi(keys[1])
					tmp := keys[0] + ":" + strconv.Itoa(vtmp-um[keys[0]])
					str = append(str, tmp)

				}
				fmt.Println(strings.Join(str, ","))

			}

		}

	}

}
