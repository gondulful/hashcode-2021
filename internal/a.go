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
	out := Output{}

	out.T2 = make([]int, 0)
	total := s.in.M

	for i:= 0; total - 2 > 0 && i < s.in.T2; i++ {
		total -= 2
		out.T2 = append(out.T2, 0)
	}

	for i:= 0; total - 3 > 0 && i < s.in.T3; i++ {
		total -= 3
		out.T3 = append(out.T3, 0)
	}

	for i:= 0; total - 4 > 0 && i < s.in.T4; i++ {
		total -= 4
		out.T4 = append(out.T4, 0)
	}

	return out
}
