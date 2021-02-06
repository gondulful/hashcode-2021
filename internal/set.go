package internal

type Set interface {
	HasElement(string) bool
	Count() int
	Intersect(Set) Set
	Union(Set) Set
}



