package main

import (
	"fmt"
	"os"
	"strconv"
	"flag"
)

type fibPred func (x int, ind int) bool

func fibWhile (aPred fibPred, bPred fibPred) [] int {
	a := 1
	b := 1
	out := []int{}
	ind := 0
	for aPred(a, ind) {
		out = append(out, a, b)
		if bPred(b, ind) { out = append(out, b) }
		a = a + b
		b = a + b
		ind += 1
	}
	return out
}

func fib (limit int) [] int {
	return fibWhile(
		func (a int, ind int) bool {
			return a <= limit
		},
		func (b int, ind int) bool {
			return b <= limit
		})
}

func fibIterations (numInts int) [] int {
	lte := func (_, ind int) bool {
		return ind <= numInts
	}

	return fibWhile(lte, lte)
}

func printNumList (list []int ) {
	for _, value := range list {
		fmt.Print(value, " ")
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("`fib` requires one flag 'n' for fib at index 'n' or 'l' for fib close-to/upto some limit")
		return
	}

	byLimit := flag.Int("l", 0, "Limit or \"upto limit\"")
	byNumber := flag.Int("n", 0, "Number or 'x' \"number\" results")

	flag.Parse();

	if *byLimit > 0 {
		printNumList(fib(*byLimit))
		return
	} else if *byNumber > 0 {
		printNumList(fibIterations(*byNumber))
		return
	}

	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Unable to convert value to `int`.  Value passed in: ", limit)
		return
	}
}
