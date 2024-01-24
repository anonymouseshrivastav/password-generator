package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: passgen <username>")
		return
	}
	name := os.Args[1]

	wordlist := ""
	wordlist = genPass(name, wordlist)

	name_2 := []rune(name)
	name_2[0] = unicode.ToUpper(name_2[0])
	name_3 := string(name_2)
	wordlist = genPass(name_3, wordlist)

	error := os.WriteFile("pass_list.txt", []byte(wordlist), 0644)

	if error != nil {
		fmt.Println(error.Error())
	}
}

func genPass(name, word string) string {

	symbols := []string{"", "!", "%", "@", "#", "$", "&"}

	for _, symbol := range symbols {
		for i := 0; i <= 1000; i++ {
			word = word + name + symbol + strconv.Itoa(i) + "\n"
		}
		for i := 1500; i <= 2100; i++ {
			word = word + name + symbol + strconv.Itoa(i) + "\n"
		}

		for _, subSymbol1 := range symbols[2:] {
			for _, subSymbol2 := range symbols[2:] {
				for _, subSymbol3 := range symbols[2:] {
					for _, subSymbol4 := range symbols[2:] {
						word = word + name + subSymbol1 + subSymbol2 + subSymbol3 + subSymbol4 + "\n"
					}
				}
			}
		}
	}

	return word
}
