package main

import (
	"fmt"
	"os"
	"strconv"
	"flag"
	fib "github.com/elycruz/go-fib/math"
)

func printNumList (list []uint64 ) {
	for _, value := range list {
		fmt.Print(value, " ")
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("`FibByLimit` requires one flag 'n' for FibByLimit at index 'n' or 'l' for FibByLimit close-to/upto some limit")
		return
	}

	byLimit := flag.Uint64("l", 0, "Prints fib numbers up to 'closest to limit' or 'up to limit' where `l` is 'limit'")
	byIterations := flag.Uint64("i", 0, "Prints fib numbers up to 'i' iterations")
	byNumber := flag.Uint64("n", 0, "Prints fib number index 'n'")

	flag.Parse();

	if *byLimit > 0 {
		printNumList(fib.FibByLimit(*byLimit))
	} else if *byNumber > 0 {
		fmt.Println(fib.NthFibNumber(*byNumber))
	} else if *byIterations > 0 {
		printNumList(fib.FibByIterations(*byNumber))
	} else {
		limit, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("Unable to convert value to `int`.  Value passed in: ", limit)
		}
		printNumList(fib.FibByLimit(limit))
	}

}
