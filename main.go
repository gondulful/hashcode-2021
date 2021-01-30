package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/a.txt`)
	a := internal.NewA(in)
	out := a.Solve()
	internal.WriteFile(out, `a_out.txt`)
}