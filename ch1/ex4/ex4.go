// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupfiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {

	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupfiles, arg)
			f.Close()
		}
	}
	for line, s := range dupfiles {
		ss := strings.Split(s, " ")
		ss = removeDuplicate1(ss)
		dupfiles[line] = strings.Join(ss, " ")
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, dupfiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupfiles map[string]string, filename string) {
	input := bufio.NewScanner(f)
	sep := " "
	for input.Scan() {
		counts[input.Text()]++
		dupfiles[input.Text()] += sep + filename
	}
	// NOTE: ignoring potential errors from input.Err()
}

func removeDuplicate1(args []string) []string {
	results := make([]string, 0, len(args))
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}

//!-
