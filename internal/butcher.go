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

type ButcherPizzeria struct {
	pizzas []Pizza
}

func NewButcherPizzeria(pizzas []Pizza) *ButcherPizzeria {
	return &ButcherPizzeria{
		pizzas: pizzas,
	}
}

func (b *ButcherPizzeria) GetPizzaWithMaxIngr() Set {
	var index int


	for i, pizza := range b.pizzas {
		if pizza.Count() > b.pizzas[index].Count() {
			index = i
		}
	}

	return b.takePizza(index)
}

func (b *ButcherPizzeria) takePizza(index int) Set {
	pizza := b.pizzas[index]

	newPizzas := make([]Pizza, 0)
	for i, p := range b.pizzas {
		if i != index {
			newPizzas = append(newPizzas, p)
		}
	}

	b.pizzas = newPizzas
	return pizza
}

func (b ButcherPizzeria) HasPizzas(i int) bool {
	return len(b.pizzas) >= i
}

func (b *ButcherPizzeria) FindUnlikePizza(set Set) Set {
	candidateIndex := 0
	for index, pizza := range b.pizzas {
		if set.Diff(pizza) > set.Diff(b.pizzas[candidateIndex]) {
			candidateIndex = index
		}
	}
	return b.takePizza(candidateIndex)
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
