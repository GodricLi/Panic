package main

import (
	"fmt"
)

func test_panic(x int) {
	//设置recover,放在出错代码的前面
	defer func() {
		// recover()
		// fmt.Println(recover())	//recover 可以打印panic的错误信息

		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var a [10]int
	a[x] = 666
	fmt.Println("end")

}

func main() {
	test_panic(12) //panic: runtime error: index out of range
}
