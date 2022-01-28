package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Print(Glass())
	fmt.Print(Hello())
	fmt.Print(Opt())
	fmt.Print(Go())
}

func Glass() string {
	return quote.Glass()
}

func Hello() string {
	return quote.Hello()
}

func Opt() string {
	return quote.Opt()
}

func Go() string {
	return quote.Go()
}
