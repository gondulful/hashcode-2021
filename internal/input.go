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

func (p Pizza) Diff(set Set) int {
	count := p.Intersect(set).Count()
	return set.Count() + p.Count() - count
}

type InputStreet struct {
	B, E int
	name string
	L int
}

type InputPath struct {
	P int
	Streets []string
}

type Input struct {
	D int // the duration of the simulation, in seconds
	I int // the number of intersections (with IDs from 0 to I -1 ),
	S int // the number of streets
	V int // the number of cars
	F int // the bonus points for each car that reachs its destination before time D
	Streets []InputStreet
	Paths []InputPath
}

func ReadFile(filename string) (i Input) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(fd)
	// first line
	sc.Scan()
	line := sc.Text()
	fields := strings.Fields(line)
	i.D = toInt(fields[0])
	i.I = toInt(fields[1])
	i.S = toInt(fields[2])
	i.V = toInt(fields[3])
	i.F = toInt(fields[4])
	// streets
	for j := 0; j < i.S; j++ {
		sc.Scan()
		line := sc.Text()
		fields := strings.Fields(line)

		var street InputStreet
		street.B = toInt(fields[0])
		street.E = toInt(fields[1])
		street.name = fields[2]
		street.L = toInt(fields[3])

		i.Streets = append(i.Streets, street)
	}
	// paths
	for j := 0; j < i.V; j++ {
		sc.Scan()
		line := sc.Text()
		fields := strings.Fields(line)

		var path InputPath
		path.P = toInt(fields[0])
		path.Streets = fields[1:]

		i.Paths = append(i.Paths, path)
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