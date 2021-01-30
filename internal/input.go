package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pizza map[string]struct{}

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
