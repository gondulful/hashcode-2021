package main

import (
	"fmt"

	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/e.txt`)
	fmt.Println(in)
	//a := internal.NewA(in)
	//out := a.Solve()
	//internal.WriteFile(out, `c_out.txt`)

	//butcher := internal.NewButcher(in, internal.NewButcherPizzeria(in.Pizzas))
	//outOut := butcher.Solve()
	//internal.WriteFile(outOut, `a_butcher_out.txt`)
}