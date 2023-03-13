package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a litter lamb"

func hasHead(s, head string) {
	fmt.Printf("string %s has the head %s: %v\n", s, head, strings.HasPrefix(s, head))
}

func hasTail(s, tail string) {
	fmt.Printf("string %s has the tail %s: %v\n", s, tail, strings.HasSuffix(s, tail))
}

func containStr(s, sub string) {
	fmt.Printf("string %s contains the substring %s: %v\n", s, sub, strings.Contains(s, sub))
}

func main() {
	hasHead(refString, "Mary")
	hasTail(refString, "lamb")
	containStr(refString, "had")
}
