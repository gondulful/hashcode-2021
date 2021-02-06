package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/d.txt`)

	butcher := internal.NewButcher(in, internal.NewButcherPizzeria(in.Pizzas))
	outOut := butcher.Solve()
	internal.WriteFile(outOut, `d_butcher_out.txt`)
}