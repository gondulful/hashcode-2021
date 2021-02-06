package internal

type Set interface {
	CountIntersect(Set) int
	Union(Set) Set
}



