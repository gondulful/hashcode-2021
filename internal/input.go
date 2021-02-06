package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pizza map[string]struct{}

func (p Pizza) Count() int {
	return len(p)
}

func (p Pizza) HasElement(s string) bool {
	_, ok := p[s]
	return ok
}

func (p Pizza) Intersect(set Set) Set {
	newPizza := make(Pizza, set.Count())
	for ingredient := range p {
		if set.HasElement(ingredient) {
			newPizza[ingredient] = struct{}{}
		}
	}

	return newPizza
}

func (p Pizza) Union(set Set) Set {
	newPizza := make(Pizza, p.Count() + set.Count())
	for k := range p {
		newPizza[k] = struct{}{}
	}
	for k := range set.(Pizza) {
		newPizza[k] = struct{}{}
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
			p := make(Pizza, len(fields[1:]))
			for _, ingredient := range fields[1:] {
				p[ingredient] = struct{}{}
			}

			i.pizzas = append(i.pizzas, p)
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
