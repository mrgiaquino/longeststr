package main

import (
	"fmt"
	"log"
	"strings"
)

func longestStr(c, subStr string) string {

	var firstIndex, lastIndex int

	//  get index of nearest c and index of farthest c

	if firstIndex = strings.Index(subStr, c); firstIndex < 0 {
		log.Fatalf("Character %s doesn't exist \n", c)
	}

	if lastIndex = strings.LastIndex(subStr, c); lastIndex == firstIndex {
		log.Fatalf("Character %s only exists once \n", c)
	}

	return subStr[firstIndex : lastIndex+1]
}

func main() {
	fmt.Println("--- longeststr ---")
	fmt.Println(longestStr("d", "SAasdDASDASJDUASHFUFUASHFUA9WQITdJHNGd"))
}
