package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	alphaLower = "abcdefghijklmnopqrstuvwxyz"
	alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers    = "0123456789"
	specials   = "!@#$%^&*"
)

//go:embed eff_large_wordlist.txt
var wordList string

var wordMap = make(map[string]string)

func initWordMap() {
	lines := strings.Split(wordList, "\n")

	for i := range len(lines) - 1 {
		parts := strings.Split(lines[i], "\t")
		key := parts[0]
		word := parts[1]
		wordMap[key] = word
	}
}

func generatePassword(n int, s []string) (p string) {
	var set string

	for _, i := range s {
		switch i {
		case "a":
			set += alphaLower
		case "A":
			set += alphaUpper
		case "0":
			set += numbers
		case "!":
			set += specials
		}
	}

	for range n {
		randIdx := rand.Intn(len(set))

		c := string(set[randIdx])

		p += c
	}

	return
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func generatePassphrase(n int, capitalize bool, includeNumber bool, separator string) (p string) {
	var words []string

	for range n {
		var key string
		for range 5 {
			key += strconv.Itoa(rand.Intn(6) + 1)
		}

		word := wordMap[key]

		if capitalize {
			words = append(words, capitalizeFirstLetter(word))
			continue
		}
		words = append(words, word)
	}

	if includeNumber {
		k := rand.Intn(n)
		r := rand.Intn(10)
		words[k] = words[k] + strconv.Itoa(r)
	}

	p = strings.Join(words, separator)

	return
}

func main() {
	n := flag.Int("n", 14, "number of characters (password) / words (passphrase)")
	s := flag.String("s", "a,A,0", "comma-separated set of characters to pick from. values: a,A,0,!")
	t := flag.String("t", "password", "type of secret. values: password, passphrase")
	c := flag.Bool("c", false, "capitalize passphrase?")
	N := flag.Bool("N", false, "include number in passphrase?")
	w := flag.String("w", " ", "word separator for passphrase")
	flag.Parse()

	switch *t {
	case "password":
		if *n < 5 {
			fmt.Println("password should be min 5 chars")
			os.Exit(1)
		}
		if *n > 128 {
			fmt.Println("password should be max 128 chars")
			os.Exit(1)
		}

		choices := strings.Split(*s, ",")
		if len(choices) < 1 {
			fmt.Println("wrong set chosen")
			os.Exit(1)
		}
		p := generatePassword(*n, choices)
		fmt.Println(p)
	case "passphrase":
		if *n < 3 {
			fmt.Println("passphrase should be min 3 words")
			os.Exit(1)
		}
		if *n > 20 {
			fmt.Println("passphrase should be max 20 words")
			os.Exit(1)
		}

		initWordMap()
		p := generatePassphrase(*n, *c, *N, *w)
		fmt.Println(p)
	}
}
