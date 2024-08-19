package main

/**
rune == character
*/

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "123ts"
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf(" %#U starts at %d \n ", runeValue, idx)
	}
	fmt.Println("using decodeRuneInString")

	for i := 0; i < len(s); i++ {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf(" %#U starts at %d width= %d \n ", runeValue, i, width)

		examineRune(runeValue)
	}

}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("fount tee")
	} else if r == 's' {
		fmt.Println("founf so sua")
	}

}
