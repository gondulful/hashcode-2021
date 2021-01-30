package internal

import (
	"os"
	"strconv"
	"strings"
)

var toString = strconv.Itoa

type Output struct {
	T2 [][]int
	T3 [][]int
	T4 [][]int
}

func WriteFile(o Output, file string) {
	fd, err := os.OpenFile(file, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	D := len(o.T2) + len(o.T3) + len(o.T4)
	_, err = fd.WriteString(toString(D))
	if err != nil {
		panic(err)
	}
	_, err = fd.Write([]byte{'\n'})
	if err != nil {
		panic(err)
	}

	for i := range o.T2 {
		fd.WriteString("2 " + strings.Join(toStringSlice(o.T2[i]), " ") + "\n")
	}
	for i := range o.T3 {
		fd.WriteString("3 " + strings.Join(toStringSlice(o.T3[i]), " ") + "\n")
	}
	for i := range o.T4 {
		fd.WriteString("4 " + strings.Join(toStringSlice(o.T4[i]), " ") + "\n")
	}
}

func toStringSlice(ints []int) []string {
	result := make([]string, len(ints))
	for i, int := range ints {
		result[i] = toString(int)
	}

	return result
}

