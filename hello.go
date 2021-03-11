package hello

import (
"fmt"
"rsc.io/quote"
)

func Hello() string {
	fmt.Println("Hello, this is from Hello ")
	return quote.Hello()
}
