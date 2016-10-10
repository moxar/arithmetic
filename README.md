# Arithmetic parser

This library parses and solves a mathematical expression.

It uses a Shunting-Yard algorithm (aka Reverse-Polish notation).

Usual mathematic constants and functions are defined in this package (see doc).
Additional can be defined using the functions `RegisterVariable` and `RegisterFunction`.

## Getting started

```go
	package main

	import(
		"fmt"
		
		"github.com/moxar/arithmetic"
	)

	func main() {
		
		v, err := arithmetic.Parse("2+3")
		if err != nil {
			// ...
		}
		
		fmt.Println(v) // print 5
	}
```

This snippet is a very simple usage example of the arithmetic package.

## Built-in

The following operators, variables and functions are supported by the package. Note that the inputs
of the Parse method are not case sensitive.

### Operators

The package supports the following operators:

	* +  (sum)
	* -  (difference)
	* /  (quotient)
	* x  (product)
	* %  (modulo). Note: modulo requires both operand to be integer
	* >  (greater)
	* >= (greater or equal)
	* == (equal)
	* <= (lower or equal)
	* <  (lower)
	* != (different)
	* && (and). Note: "and" is an alias to "&&"
	* || (or). Note: "or" is an alias to "||"
	* !  (not). Note: "not" is an alias to "!"
	
### Variables

The package supports the following variables from the [math](https://golang.org/pkg/math/) package:

	* e
	* pi
	* phi
	
	* sqrt2
	* sqrte
	* sqrtpi
	* sqrtphi
	
	* ln2
	* ln2e
	* ln10
	* log10e
	
Additionnaly, it supports the boolean `true` and `false`.

### Functions

The package supports the following functions:

	* max(a, b, c...). Note: Max only supports float, but can use any number of arguments.
	* min(a, b, c...). Note: Min only supports float, but can use any number of arguments.
	* mean(a, b, c...). Note: Mean only supports float, but can use any number of arguments.
	* if(condition, success, failure).

## Extension

The package can be extended. You can add variable, functions or expression to be automaticaly recognized by the parser. The extension is made possible by the `RegisterXXX` functions, which works in a similar way than the sql drivers. The functions will panic if there is a conflict between the name of the newly defined variable/function and an existing one.

### Variable

A variable is a mapping from a string to a value. In this example, when the parser detects
"dayInYear", it converts it to 365.

```go
	package main
	
	import "github.com/moxar/arithmetic"

	func init() {
		arithmetic.RegisterVariable("dayInYear", 365)
	}
	
	func main() {
		
		v, err := arithmetic.Parse("dayInYear * 2")
		if err != nil {
			// ...
		}
		
		fmt.Println(v) 
		
		// Output: 730
	}
```

### Function

A function is a mapping from a string to a function that returns a value. In this example, when the parser detects "contains", it checks if the second argument is contained within the first one.

```go
	package main
	
	import (
		"fmt"
	
		"github.com/moxar/arithmetic"
	)

	func init() {

		// Register a new function, that increments the value by one.
		arithmetic.RegisterFunction("increment", func(args ...interface{}) (interface{}, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("increment requires one argument, %d provided", len(args))
			}

			f, ok := arithmetic.ToFloat(args[0])
			if !ok {
				return nil, fmt.Errorf("increment requires integer argument, %v (%T) provided", args[0], args[0])
			}

			return int(f) + 1, nil
		})
	}
	
	func main() {

		v, err := Parse("increment(2)")
		if err != nil {
			// ...
		}

		fmt.Println(v)
		// Output: 3
	}
```

### Expression

An expression is a function that transforms an input to a value. The boolean tells whether the input could be transformed or not. Because the expressions are tested in an undefined order, you have to ensure there is no overlap between them.

```go
	package main
	
	import (
		"fmt"
		"strings"
	
		"github.com/moxar/arithmetic"
	)

	func init() {
		arithmetic.RegisterExpression(foo)
	}
	
	func foo(input string) (interface{}, bool) {
		if len(input) == 0 {
			return nil, false
		}

		if input[0] != 'f' {
			return nil, false
		}
		

		for i, r := range input {
		
			if r != 'o' {
				return nil, false
			}
		}

		return true, true
	})
	
	func main() {

		v, err := arithmetic.Parse("fooOOoOoooOOOooOo && true")
		if err != nil {
			// ...
		}

		fmt.Println(v) 
		// Output: true
	}
```

## Disclaimer

This lib is still at early development stages. The API is most likely susceptible to change.
