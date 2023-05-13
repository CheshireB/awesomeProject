package pair_usage_2

type pair interface {
	Do()
}

func SomeUsage(pairs ...pair) []pair {
	return pairs
}
