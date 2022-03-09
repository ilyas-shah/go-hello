package main

import (
	"fmt"

	"github.com/ilyas-shah/go-hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("hello"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
