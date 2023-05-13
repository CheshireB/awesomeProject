package pair_usage

type pair interface {
	Do()
}

func SomeUsage(pairs ...pair) {
	for _, p := range pairs {
		p.Do()
	}
}
