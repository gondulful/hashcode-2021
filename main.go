package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/a.txt`)
	//a := internal.NewA(in)
	//out := a.Solve()
	//internal.WriteFile(out, `c_out.txt`)

	butcher := internal.NewButcher(in, internal.NewButcherPizzeria(in.Pizzas))
	outOut := butcher.Solve()
	internal.WriteFile(outOut, `a_butcher_out.txt`)
}