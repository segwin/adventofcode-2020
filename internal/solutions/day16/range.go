package day16

type Range struct {
	Min int
	Max int
}

func (r *Range) Contains(value int) bool {
	return value >= r.Min && value <= r.Max
}
