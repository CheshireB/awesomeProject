package main

import (
	"awesomeProject/pair"
	"awesomeProject/pair_usage"
)

func main() {
	p1 := pair.NewBuilder(1)
	p2 := pair.NewBuilder(2)

	pair_usage.SomeUsage(p1, p2)
}
