package main

import "fmt"

func main() {
	// Tạo hình chữ nhật
	s := []Student{
		{
			Person: Person{
				name: "An",
				age:  18,
			},
			class: "12A",
		},
		{
			Person: Person{
				name: "An",
				age:  18,
			},
			class: "12A",
		},
		{
			Person: Person{
				name: "An",
				age:  18,
			},
			class: "12A",
		},
	}

	// Gọi method Area()
	for i, student := range s {
		fmt.Printf("%d. Tên: %s, Tuổi: %d, Lớp: %s\n", i+1, student.name, student.age, student.class)
	}
}
