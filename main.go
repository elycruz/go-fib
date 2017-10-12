package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib (limit int) [] int {
	a := 1
	b := 1
	out := []int{}
	for a <= limit {
		out = append(out, a)
		if b <= limit { out = append(out, b) }
		a = a + b
		b = a + b
	}
	return out
}

func printNumList (list []int ) {
	for _, value := range list {
		fmt.Print(value, " ")
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("`fib` requires one argument of type int")
		return
	}
	
	limit, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Unable to convert value to `int`.  Value passed in: ", limit)
		return
	}
	printNumList(fib(limit))
}

