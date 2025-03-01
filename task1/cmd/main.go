package main

import (
	"challenges/bigsum"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout) // Ghi log ra console
	b := bigsum.BigNumber{}

	for {
		var x, y string

		fmt.Print("Nhập số thứ nhất: ")
		fmt.Scanln(&x)

		fmt.Print("Nhập số thứ hai: ")
		fmt.Scanln(&y)

		result := b.Sum(x, y)
		fmt.Printf("Tổng của %s và %s là: %s\n", x, y, result)

		var choice string
		fmt.Print("Bạn có muốn tiếp tục? (y/n): ")
		fmt.Scanln(&choice)
		if choice != "y" {
			fmt.Println("Chương trình kết thúc.")
			break
		}
	}
}
