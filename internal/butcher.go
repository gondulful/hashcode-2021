package internal

type butcher struct {
	in Input
	p Pizzeria
}

type Pizzeria interface {
	GetPizzaWithMaxIngr() Set
	HasPizzas(int) bool
	FindUnlikePizza(Set) Set
}

func NewButcher(in Input, p Pizzeria) Solver {
	return &butcher{
		in,
		p,
	}
}

func (s *butcher) Solve() Output {
	out := Output{}

	for i := 0; s.p.HasPizzas(4) && i < s.in.T4; i++ {
		first := s.p.GetPizzaWithMaxIngr()
		second := s.p.FindUnlikePizza(first)
		wantedThird := first.Union(second)
		third := s.p.FindUnlikePizza(wantedThird)
		wantedFourth := wantedThird.Union(third)
		fourth := s.p.FindUnlikePizza(wantedFourth)

		out.T4 = append(out.T4, []int{
			first.Id(),
			second.Id(),
			third.Id(),
			fourth.Id(),
		})
	}

	for i := 0; s.p.HasPizzas(3) && i < s.in.T3; i++ {
		first := s.p.GetPizzaWithMaxIngr()
		second := s.p.FindUnlikePizza(first)
		wantedThird := first.Union(second)
		third := s.p.FindUnlikePizza(wantedThird)

		out.T3 = append(out.T3, []int{
			first.Id(),
			second.Id(),
			third.Id(),
		})
	}

	for i := 0; s.p.HasPizzas(2) && i < s.in.T2; i++ {
		first := s.p.GetPizzaWithMaxIngr()
		second := s.p.FindUnlikePizza(first)

		out.T2 = append(out.T2, []int{
			first.Id(),
			second.Id(),
		})
	}

	return out
}

