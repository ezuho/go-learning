package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// 这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
	fmt.Println("After mayPanic()")
}
