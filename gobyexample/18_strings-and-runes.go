package main

import "unicode/utf8"
import "fmt"

func main() {
	const s = "aaa"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
	}

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRUNHEiNString")
	for i, w := 0, 0; i < len(s); i+= w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width 

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'a' {
		fmt.Println("found so sua")
	}
}
