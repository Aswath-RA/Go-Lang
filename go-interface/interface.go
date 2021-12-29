package main

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' || rune == 'A' || rune == 'E' || rune == 'I' || rune == 'O' || rune == 'U' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {

	var name string
	fmt.Print("Enter a String to find vowel in it: ")
	fmt.Scan(&name)

	var v VowelsFinder

	v = MyString(name)
	fmt.Printf("Vowels are %c", v.FindVowels())

}
