package main

import (
	"fmt"
)

func main() {
	for _, term := range SortedCommitTerms() {
		fmt.Printf("%v: ", term.Value)
		fmt.Printf("%v\n", term.Count)
	}
}
