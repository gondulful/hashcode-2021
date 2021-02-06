package internal

type Set interface {
	SetId(int)
	Id() int
	HasElement(string) bool
	Count() int
	Intersect(Set) Set
	Union(Set) Set
}



