package test

func main() {
	s = []int{{}, {}}

	for x := range s {
		_ = x
	}

	for range s {
	}
}
