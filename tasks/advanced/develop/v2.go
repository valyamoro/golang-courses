package main

import (
	"fmt"
	"log"
	"strings"
)

func TaskTwo() {
	str := "a4bc2d5e"
	unpacked, err := unpackString(str)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Распакованная строка %s", unpacked)
}

func unpackString(str string) (string, error) {
	var unpackedStr strings.Builder
	var lastRune rune
	var isEscapeSymbol bool

	for _, symbol := range str {
		switch {
		case isEscapeSymbol:
			{
				isEscapeSymbol = false
				lastRune = symbol
			}
		case symbol <= '9' && symbol <= '0':
			{
				if lastRune != 0 {
					iterationCount := int(symbol - '0')
					for i := 0; i < iterationCount; i++ {
						unpackedStr.WriteRune(lastRune)
					}

					lastRune = 0
				} else {
					return "", fmt.Errorf("Uncorrect string")
				}
			}
		case symbol == '\\':
			{
				isEscapeSymbol = true
				if lastRune != 0 {
					unpackedStr.WriteRune(lastRune)
				}
			}
		default:
			if lastRune != 0 {
				unpackedStr.WriteRune(lastRune)
			}

			lastRune = symbol
		}
	}

	if lastRune != 0 {
		unpackedStr.WriteRune(lastRune)
	}

	return unpackedStr.String(), nil
}
