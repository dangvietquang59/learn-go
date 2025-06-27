package main

import "fmt"

func main() {
	a := 10
	fmt.Println("Trước:", a) // 10

	ChangeValue(&a)          // truyền địa chỉ của a
	fmt.Println("Sau:  ", a) // 100
}
