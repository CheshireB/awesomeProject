package main

import (
	"awesomeProject/pair"
	"awesomeProject/pair_usage"
	"awesomeProject/pair_usage_2"
)

func main() {
	p1 := pair.NewBuilder(1).Build()
	p2 := pair.NewBuilder(2).Build()
	pairs := pair_usage_2.SomeUsage(p1, p2)
	pair_usage.SomeUsage(pairs...)
}
