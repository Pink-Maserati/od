package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {

	// m := make(map[int]int, 10)
	// for i := 1; i <= 10; i++ {
	// 	m[i] = i
	// }

	// for k, v := range m {
	// 	//	fmt.Println("hahhha  k ->", k, "v ->", v)
	// 	go func() {
	// 		fmt.Println("k ->", k, "v ->", v)
	// 	}()
	// }
	// time.Sleep(1 * time.Second)

	// a := 'a'
	// fmt.Println(string(a))
	num1 := "11"
	num2 := "123"
	fmt.Println(addStrings(num1, num2))
	a := big.NewInt(123)
	b := big.NewInt(234)
	fmt.Println(a.Add(a, b))
}

func addStrings(num1 string, num2 string) string {
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	add := 0
	res := ""
	for n1 >= 0 || n2 >= 0 || add != 0 {

		tmp1 := 0
		tmp2 := 0
		if n1 >= 0 {
			tmp1, _ = strconv.Atoi(string(num1[n1]))
		}
		if n2 >= 0 {
			tmp2, _ = strconv.Atoi(string(num2[n2]))
		}

		tmp := tmp1 + tmp2 + add
		res = strconv.Itoa(tmp%10) + res
		add = tmp / 10
		n1--
		n2--

	}
	return res

}
