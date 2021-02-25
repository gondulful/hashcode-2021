package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type InputStreet struct {
	B, E int
	Name string
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
		street.Name = fields[2]
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