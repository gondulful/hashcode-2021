package internal

import (
	"sync"
)

type butcher struct {
	in Input
	p Pizzeria
}

type Pizzeria interface {
	GetPizzaWithMaxIngr() Set
	HasPizzas(int) bool
	CountPizzas() int
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

func (b ButcherPizzeria) CountPizzas() int {
	return len(b.pizzas)
}

func (b *ButcherPizzeria) FindUnlikePizza(set Set) Set {
	//t := time.Now()
	const workers = 16
	var wg sync.WaitGroup
	limit := len(b.pizzas) / workers
	var results [workers]int
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(worker int) {
			offset := worker * limit
			candidateIndex := offset
			bestDiff := set.Diff(b.pizzas[candidateIndex])
			for i := offset; i < offset + limit; i++ {
				diff := set.Diff(b.pizzas[i])
				if diff > bestDiff {
					candidateIndex = i
					bestDiff = diff
				}
			}
			results[worker] = candidateIndex
			wg.Done()
		}(i)
	}
	wg.Wait()

	bestDiff := set.Diff(b.pizzas[results[0]])
	trueCandidate := 0
	for _, candidateIndex := range results {
		diff := set.Diff(b.pizzas[candidateIndex])
		if diff > bestDiff {
			bestDiff = diff
			trueCandidate = candidateIndex
		}
	}


	//fmt.Println("find pizza:", time.Since(t))
	//t = time.Now()
	p := b.takePizza(trueCandidate)
	//fmt.Println("take pizza:", time.Since(t))
	return p
}

func NewButcher(in Input, p Pizzeria) Solver {
	return &butcher{
		in,
		p,
	}
}

func (s *butcher) Solve() Output {
	out := Output{}

	//for i := 0; s.p.HasPizzas(4) && i < s.in.T4; i++ {
	//	fmt.Printf("Pizza4: %d\n", s.p.CountPizzas())
	//
	//	first := s.p.GetPizzaWithMaxIngr()
	//
	//	//t := time.Now()
	//	second := s.p.FindUnlikePizza(first)
	//	//fmt.Println("FindUnlikePizza:", time.Since(t))
	//	wantedThird := first.Union(second)
	//	third := s.p.FindUnlikePizza(wantedThird)
	//	wantedFourth := wantedThird.Union(third)
	//	fourth := s.p.FindUnlikePizza(wantedFourth)
	//
	//	out.T4 = append(out.T4, []int{
	//		first.Id(),
	//		second.Id(),
	//		third.Id(),
	//		fourth.Id(),
	//	})
	//}
	//
	//for i := 0; s.p.HasPizzas(3) && i < s.in.T3; i++ {
	//	fmt.Printf("Pizza3: %d\n", s.p.CountPizzas())
	//	first := s.p.GetPizzaWithMaxIngr()
	//	second := s.p.FindUnlikePizza(first)
	//	wantedThird := first.Union(second)
	//	third := s.p.FindUnlikePizza(wantedThird)
	//
	//	out.T3 = append(out.T3, []int{
	//		first.Id(),
	//		second.Id(),
	//		third.Id(),
	//	})
	//}
	//
	//for i := 0; s.p.HasPizzas(2) && i < s.in.T2; i++ {
	//	fmt.Printf("Pizza2: %d\n", s.p.CountPizzas())
	//	first := s.p.GetPizzaWithMaxIngr()
	//	second := s.p.FindUnlikePizza(first)
	//
	//	out.T2 = append(out.T2, []int{
	//		first.Id(),
	//		second.Id(),
	//	})
	//}

	return out
}

