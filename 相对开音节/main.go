package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 相对开音节

// 相对开音节构成的结构为辅音+元音（aeiou）+辅音(r除外)+e，常见的单词有bike、cake等。

// 给定一个字符串，以空格为分隔符，反转每个单词中的字母，若单词中包含如数字等其他非字母时不进行反转。

// 反转后计算其中含有相对开音节结构的子串个数（连续的子串中部分字符可以重复）。

// 输入描述:

// 字符串，以空格分割的多个单词，字符串长度<10000，字母只考虑小写

// 输出描述:

// 含有相对开音节结构的子串个数，注：个数<10000

// 示例1
// 输入

// ekam a ekac
// 输出

// 2
// 说明

// 反转后为 make a cake 其中make、cake为相对开音节子串，返回2

// 示例2

// 输入

// !ekam a ekekac
// 输出

// 2
// 说明
// 反转后为!ekam a cakeke 因!ekam含非英文字符所以未反转，其中 cake、keke为相对开音节子串，返回2

// 输入为一行字符串，定义一个名词叫音节：一个辅音,一个元音(aeiou),一个辅音（除了r），一个e。
// 如果字符串中的单词只有字母，则单词反转。
// 输出最多有几个音符。
func judgeRevserse(str string) bool {
	//正则匹配吧
	r, _ := regexp.Compile("[a-z]+")
	s := r.ReplaceAllString(str, "")
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}

func reverseString(str string) string {
	tmp := []byte(str)
	left := 0
	right := len(tmp) - 1
	for left < right {
		tmp[left], tmp[right] = tmp[right], tmp[left]
		left++
		right--
	}

	return string(tmp)
}
func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'

}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	txts := strings.Split(input, " ")
	count := 0
	for i := 0; i < len(txts); i++ {
		//需不需要反转
		tmp := txts[i]
		if judgeRevserse(tmp) {
			//反转吧
			tmp = reverseString(tmp)
		}
		if len(tmp) < 4 {
			continue
		}

		//开音节
		for j := 0; j < len(tmp)-3; j++ {
			if !isVowel(tmp[j]) && isVowel(tmp[j+1]) && !isVowel(tmp[j+2]) && tmp[j+2] != 'r' && tmp[j+3] == 'e' {
				count++
			}

		}

	}
	fmt.Println(count)

}
