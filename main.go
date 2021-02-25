package main

import (
	"hashcode-2021/internal"
)

func main() {
	in := internal.ReadFile(`files/e.txt`)
	_ = in

	counts := make(map[string]int)
	for _, p := range in.Paths {
		for _, s := range p.Streets {
			counts[s]++
		}
	}

	outm := make(map[int]internal.OutputIntersection)
	for _, s := range in.Streets {
		if cnt := counts[s.Name]; cnt == 0 {
			continue
		}

		if _, ok := outm[s.E]; !ok {
			outm[s.E] = internal.OutputIntersection{Id: s.E}
		}

		sl := outm[s.E]
		sl.StreetLights = append(sl.StreetLights, internal.OutputStreetLights{Name:s.Name, Sec: 1})
		outm[s.E] = sl
	}

	intersects := make([]internal.OutputIntersection, 0, len(outm))
	for _, intersect := range outm {
		intersects = append(intersects, intersect)
	}

	internal.WriteFile(internal.Output{Intersections: intersects}, `a_out.txt`)
	//a := internal.NewA(in)
	//out := a.Solve()
	//internal.WriteFile(out, `c_out.txt`)

	//butcher := internal.NewButcher(in, internal.NewButcherPizzeria(in.Pizzas))
	//outOut := butcher.Solve()
	//internal.WriteFile(outOut, `a_butcher_out.txt`)
}