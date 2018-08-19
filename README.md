# go-fib
Basic fibonacci sequence generator.

## Usage
`$ go-fib {number}` - Fibonacci up to or near to 'limit' value.
`$ go-fib -l {number}` - Same as `$go-fib {number}`.
`$ go-fib -i {number}` - Fibonacci up to 'n' iterations.
`$ go-fib -n {number}` - Nth Fibonacci number - Fib num at index `n`.

### Example output:

**Up to or near to given limit ('-l''**:
```
$ go-fib 99
0 1 1 2 3 5 8 13 21 34 55 89

$ go-fib -l 99
0 1 1 2 3 5 8 13 21 34 55 89
```

**Up to a number of given iterations ('-i')**:
```
$ go-fib -i 12
0 1 1 2 3 5 8 13 21 34 55 89
```

**Fib num at given index ('-n')**:
```
$ go-fib -i 12
89

$ go-fib -i 12
9223372036854775808
```

## Importing Fib funcs into other project
See './math/Fib.go' for exported.

## Resources
- Binet's Formula for nth Fib num: http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibFormula.html

## License
BSD 3
