package main

import (
	"fmt"
)

func main() {
	//	var a int = 10
	//	for a < 20 {
	//		fmt.Println(a)
	//		a++
	//		if a > 15 {
	//			break
	//			//continue
	//		}
	//	}

	//k1:
	//	//	fmt.Println("-----")
	//	for a < 20 {
	//		Println(a)
	//		a++
	//		if a > 15 {
	//			//break k1
	//			break k1
	//		}

	//		//		fmt.Println("hello.....world")
	//	}
	//	Println("end.........")
J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 6 {
				break J //现在终止的是j 循环，而不是i的那个
			}
			fmt.Println(i)
		}
	}

}
