package test

const ()

func main() {
	s = []int{int{}, int{}}

	for x, _ := range s {
		_ = x
	}

	for _ = range s {
	}
}
