package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Print("Nhập số phần tử: ")
	fmt.Fscan(reader, &n)

	// Tạo slice với capacity n
	arr := make([]int, 0, n)

	// Nhập từng phần tử
	for i := 0; i < n; i++ {
		var x int
		fmt.Printf("Phần tử %d: ", i+1)
		fmt.Fscan(reader, &x)
		arr = append(arr, x)
	}

	// Tính tổng
	total := CalculateTotal(arr)
	fmt.Printf("Tổng các phần tử = %d\n", total)

}
