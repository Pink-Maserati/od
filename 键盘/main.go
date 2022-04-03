package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
    有一个特殊的五键键盘
    上面有A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A
    A键在屏幕上输出一个字母A
    Ctrl-C将当前所选的字母复制到剪贴板
    Ctrl-X将当前选择的字母复制到剪贴板并清空所选择的字母
    Ctrl-V将当前剪贴板的字母输出到屏幕
    Ctrl-A选择当前屏幕中所有字母
    注意：
      1. 剪贴板初始为空
      2. 新的内容复制到剪贴板会覆盖原有内容
      3. 当屏幕中没有字母时,Ctrl-A无效
      4. 当没有选择字母时Ctrl-C、Ctrl-X无效
      5. 当有字母被选择时A和Ctrl-V这两个输出功能的键,
         会先清空所选的字母再进行输出

    给定一系列键盘输入,
    输出最终屏幕上字母的数量

    输入描述:
       输入为一行
       为简化解析用数字12345分别代替A、Ctrl-C、Ctrl-X、Ctrl-V、Ctrl-A的输入
       数字用空格分割

    输出描述:
        输出一个数字为屏幕上字母的总数量

    示例一:
        输入:
          1 1 1
        输出:
		  3
		说明:
		连续键入3个a，故屏幕上字母的长度为3

   示例二:
        输入:
          1 1 5 1 5 2 4 4
        输出:
		  2
		说明:
        输入两个a后ctrl-a选择这两个a，再输入a时选择的两个a先被清空，所以此时屏幕只有一个a，后续的ctrl-a，ctrl-c选择并复制了这一个a，最后两个ctrl-v在屏幕上输出两个a，故屏幕上字母的长度为2（第一个ctrl-v清空了屏幕上的那个a）

*/

func f1(str []string) {
	ret := strings.Builder{}
	choose := ""
	tab := ""

	for _, v := range str {
		if v == "1" {
			if len(choose) > 0 {
				ret = strings.Builder{}
				choose = ""

			}
			ret.WriteString("A")

		} else if v == "2" {
			if len(choose) > 0 {
				tab = choose
			}
		} else if v == "3" {
			if len(choose) > 0 {
				tab = choose
				choose = ""
				ret = strings.Builder{}
			}

		} else if v == "4" {

			if len(choose) > 0 {
				ret = strings.Builder{}
				choose = ""

			}
			ret.WriteString(tab)
		} else if v == "5" {
			if len(ret.String()) > 0 {
				choose = ret.String()
			}

		}

		fmt.Println(ret.String())
		//	fmt.Println("tab:", tab)
		//fmt.Println("choose:", choose)
		fmt.Println(len(ret.String()))

	}

	fmt.Println("结果：")
	fmt.Println(len(ret.String()))

}

func f2(str []string) {
	count := 0
	copyCount := 0
	selCount := 0
	for _, v := range str {
		if v == "1" {

			if selCount > 0 {
				count = 1
			} else {
				count = count + 1
			}
			selCount = 0
		} else if v == "2" {
			//复制
			copyCount = selCount

		} else if v == "3" {
			//剪贴
			count -= selCount
			copyCount = selCount
			selCount = 0
		} else if v == "4" {
			//粘贴
			count -= selCount
			selCount = 0
			count += copyCount
		} else if v == "5" {
			//全选
			selCount = count
		}
	}
	fmt.Println("f2:", count)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	str := strings.Split(input, " ")
	f1(str)
	f2(str)

}
