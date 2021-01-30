package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/e.txt`)
	a := internal.NewA(in)
	out := a.Solve()
	internal.WriteFile(out, `e_out.txt`)
}