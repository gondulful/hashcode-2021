package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/e.txt`)
	_ = in


	var out = internal.Output{
		Intersections: []internal.OutputIntersection{
			{Id: 1, StreetLights: []internal.OutputStreetLights{
				{"rue-d-athenes", 2},
				{"rue-d-amsterdam", 1},
			}},
			{Id: 0, StreetLights: []internal.OutputStreetLights{
				{"rue-de-londres", 2},
			}},
			{Id: 2, StreetLights: []internal.OutputStreetLights{
				{"rue-de-moscou", 1},
			}},
		},
	}
	internal.WriteFile(out, `a_out.txt`)
	//a := internal.NewA(in)
	//out := a.Solve()
	//internal.WriteFile(out, `c_out.txt`)

	//butcher := internal.NewButcher(in, internal.NewButcherPizzeria(in.Pizzas))
	//outOut := butcher.Solve()
	//internal.WriteFile(outOut, `a_butcher_out.txt`)
}