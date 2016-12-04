// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	num := 1000000
	start1 := time.Now()
	for i := 0; i < num; i++ {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}

	fmt.Printf("%.2f sec using strings.join\n", time.Since(start1).Seconds())

	start2 := time.Now()
	for i := 0; i < num; i++ {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}

	fmt.Printf("%.2f sec using xxx\n", time.Since(start2).Seconds())
}

//!-
