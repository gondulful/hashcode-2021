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
	total := s.in.M
	pizza := 0
	for i := 0; total - 2 >= 0 && i < s.in.T2; i++ {
		out.T2 = append(out.T2, []int{pizza , pizza + 1})
		pizza += 2
		total -= 2
	}

	for i := 0; total - 3 >= 0 && i < s.in.T3; i++ {
		out.T3 = append(out.T3, []int{pizza, pizza + 1, pizza + 2})
		pizza += 3
		total -= 3
	}

	for i := 0; total - 4 >= 0 && i < s.in.T4; i++ {
		out.T4 = append(out.T4, []int{pizza, pizza + 1, pizza + 2, pizza + 3})
		pizza += 4
		total -= 4
	}

	return out
}
