package main

import "fmt"

func main() {
	var m = map[int]string{
		0: "0",
		1: "1",
	}
	mm := m
	fmt.Printf("指针地址： %p\n", m)       //0xc42002a028 指针地址
	fmt.Printf("存指针地址的地址: %p\n", &m)  //0xc42002a028 存指针地址的地址
	fmt.Printf("存指针地址的地址: %p\n", &mm) //0xc42002a030 存指针地址的地址
	fmt.Println(m)                    // map[0:0 1:1]
	fmt.Println(mm)                   //map[1:1 0:0]
	changeMap(m)
	fmt.Printf("指针地址: %p\n", m)       //指针地址
	fmt.Printf("存指针地址的地址: %p\n", &m)  //0xc42002a028 存指针地址的地址
	fmt.Printf("存指针地址的地址: %p\n", &mm) //0xc42002a030 存指针地址的地址
	fmt.Println(m)                    //map[2:2 0:0 1:1]
	fmt.Println(mm)                   //map[0:0 1:1 2:2]

	var p *int
	var i int = 0
	p = &i
	fmt.Println(p)

	var p2 *map[int]string = &m
	fmt.Printf("%p", p2)
}
func changeMap(mmm map[int]string) {
	mmm[2] = "2"
	fmt.Println()
	fmt.Printf("指针地址: changeMap func %p\n", mmm)      //changeMap func 0xc420014150 指针地址
	fmt.Printf("存指针地址的地址: changeMap func %p\n", &mmm) //changeMap func 0xc420014150 存指针地址的地址
	fmt.Println()
}
