package main

import (
	"fmt"

	"github.com/kalifun/ggCode/pkg/font"
)

func main() {
	a := "test"
	b := font.NewColorFont(a)
	fmt.Println(b.Green())
	fmt.Println(b.Red())
	fmt.Println(b.Yellow())
	fmt.Println(b.Cyan())
	fmt.Println(b.White())
}
