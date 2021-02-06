package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pizza struct{
	id int
	ingredients map[string]struct{}
}

func (p Pizza) SetId(id int) {
	p.id = id
}

func (p Pizza) Id() int {
	return p.id
}

func (p Pizza) Count() int {
	return len(p.ingredients)
}

func (p Pizza) HasElement(s string) bool {
	_, ok := p.ingredients[s]
	return ok
}

func (p Pizza) Intersect(set Set) Set {
	newPizza := Pizza{
		id:          -1,
		ingredients: make(map[string]struct{}, set.Count()),
	}
	for ingredient := range p.ingredients {
		if set.HasElement(ingredient) {
			newPizza.ingredients[ingredient] = struct{}{}
		}
	}

	return newPizza
}

func (p Pizza) Union(set Set) Set {
	newPizza := Pizza{
		id:          -1,
		ingredients: make(map[string]struct{}, p.Count() + set.Count()),
	}
	for k := range p.ingredients {
		newPizza.ingredients[k] = struct{}{}
	}
	for k := range set.(Pizza).ingredients {
		newPizza.ingredients[k] = struct{}{}
	}

	return newPizza
}

type Input struct {
	M int
	T2 int
	T3 int
	T4 int
	pizzas []Pizza
}

func ReadFile(filename string) (i Input) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}


	var isFirstLine = true
	sc := bufio.NewScanner(fd)
	var pizzaNumber int
	for sc.Scan() {
		line := sc.Text()
		fields := strings.Fields(line)

		if isFirstLine {
			i.M = toInt(fields[0])
			i.T2 = toInt(fields[1])
			i.T3 = toInt(fields[2])
			i.T4 = toInt(fields[3])
			isFirstLine = false
		} else {
			p := Pizza{
				id:          pizzaNumber,
				ingredients: make(map[string]struct{}, len(fields[1:])),
			}
			for _, ingredient := range fields[1:] {
				p.ingredients[ingredient] = struct{}{}
			}

			i.pizzas = append(i.pizzas, p)
			pizzaNumber++
		}
	}

	err = sc.Err()
	if err != nil {
		panic(err)
	}

	return
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
