package main

import (
	"fmt"
	"log"
	"strings"
)

const lowerCase = "abcdefghijklmnopqrstuvwxyz"
const upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	ret := ""
	for _, char := range s {
		switch {
		case strings.ContainsRune(lowerCase, char):
			ret = ret + string(rotate(char, k, []rune(lowerCase)))
		case strings.ContainsRune(upperCase, char):
			ret = ret + string(rotate(char, k, []rune(upperCase)))
		default:
			ret = ret + string(char)
		}
	}
	return ret
}

func rotate(s rune, delta int32, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		log.Printf("Could not find %v in \n\t%v", s, key)
		panic("See Log")
	}

	idx = (idx + int(delta)) % len(key)
	return key[idx]
}

func main() {
	var length, delta int32
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	fmt.Printf("%d \n%s \n%d \n", length, input, delta)
	fmt.Printf("result:\n%s\n", caesarCipher(input, delta))
}
