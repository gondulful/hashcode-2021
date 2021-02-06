package internal

type Set interface {
	HasElement(string) bool
	Count() int
	CountIntersect(Set) int
	Union(Set) Set
}



