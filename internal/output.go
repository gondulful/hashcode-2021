package internal

import (
	"os"
	"strconv"
)

var toString = strconv.Itoa

type OutputStreetLights struct {
	Name string
	Sec  int
}
type OutputIntersection struct {
	Id int
	StreetLights []OutputStreetLights
}
type Output struct {
	Intersections []OutputIntersection
}

func WriteFile(o Output, file string) {
	fd, err := os.OpenFile(file, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	var numIntersections = len(o.Intersections)
	_, err = fd.WriteString(toString(numIntersections))
	if err != nil {
		panic(err)
	}
	_, err = fd.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}

	for _, intersection := range o.Intersections {
		fd.WriteString(toString(intersection.Id) + "\n")
		fd.WriteString(toString(len(intersection.StreetLights)) + "\n")
		for _, item := range intersection.StreetLights {
			fd.WriteString(item.Name + " " + toString(item.Sec) + "\n")
		}
	}
}

func toStringSlice(ints []int) []string {
	result := make([]string, len(ints))
	for i, int := range ints {
		result[i] = toString(int)
	}

	return result
}

