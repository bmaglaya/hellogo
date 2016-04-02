package main

import (
		"fmt"
		
		"github.com/bmaglaya/stringutil"
)

func main() {
	str := "Reverse me\n"
	fmt.Printf(str)
	fmt.Printf(stringutil.Reverse(str))
}