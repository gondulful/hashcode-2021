package main

import (
	"fmt"

	"./internal"
)

func main() {
	in := internal.ReadFile(`files/a.txt`)
	fmt.Println(in)
}