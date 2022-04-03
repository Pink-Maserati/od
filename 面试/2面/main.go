package main

import (
	"fmt"
	"strings"
)

func f(str string) bool {
	if len(str) == 0 || len(str) == 1 {
		return true
	}
	//fmt.Println("str:",str)
	sb := strings.Builder{}
	for i := 0; i < len(str); i++ {

		if (str[i] >= '0' && str[i] <= '9') || (str[i] >= 'a' && str[i] <= 'z') {
			sb.WriteByte(str[i])
		} else if str[i] >= 'A' && str[i] <= 'Z' {
			tmp := string(str[i])

			sb.WriteString(strings.ToLower(tmp))
		}
		//fmt.Println("sb:",sb.String())
	}
	s := sb.String()
	//fmt.Println(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}

	}
	return true

}

func main() {
	//a:="A man, a plan, a canal: Panama"
	a := "race a car"
	fmt.Println(f(a))

}
