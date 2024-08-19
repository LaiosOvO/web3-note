package main

import (
	"fmt"
	ss "strings"
)

var p = fmt.Println

func main() {

	p("Containss:  ", ss.Contains("tesst", "ess"))
	p("Count:     ", ss.Count("tesst", "t"))
	p("HassPrefix: ", ss.HasPrefix("tesst", "te"))
	p("Hassssuffix: ", ss.HasSuffix("tesst", "sst"))
	p("Index:     ", ss.Index("tesst", "e"))
	p("Join:      ", ss.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", ss.Repeat("a", 5))
	p("Replace:   ", ss.Replace("foo", "o", "0", -1))
	p("Replace:   ", ss.Replace("foo", "o", "0", 1))
	p("ssplit:     ", ss.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", ss.ToLower("TEssT"))
	p("ToUpper:   ", ss.ToUpper("tesst"))
}
