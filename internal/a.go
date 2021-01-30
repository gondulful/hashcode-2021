package internal

type a struct {
	in Input
}

func NewA(in Input) Solver {
	return &a{
		in,
	}
}

func (s *a) Solve() Output {
	return Output{
		T2: []int{1, 2, 3},
		T3: []int{1, 2, 3},
		T4: []int{1, 2, 3},
	}
}
