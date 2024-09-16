package goReloded

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func StringConvert(str string) string {
	w := ExtraSpaces(str)
	words := strings.Split(w, " ")

	for len(words) > 0 && (words[0] == "(hex)" || words[0] == "(bin)" || words[0] == "(up)" || words[0] == "(low)" || words[0] == "(cap)") {
		log.Fatalf("nothing before () to convert recheck  u're information")
	}

	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && words[i+1] == "(hex)" {
			words[i] = HexToDec(words[i])
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
		if i+1 < len(words) && words[i+1] == "(bin)" {
			words[i] = BinToDec(words[i])
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
		if i+1 < len(words) && words[i+1] == "(up)" {
			words[i] = strings.ToUpper(words[i])
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
		if i+1 < len(words) && words[i+1] == "(low)" {
			words[i] = strings.ToLower(words[i])
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
		if i+1 < len(words) && words[i+1] == "(cap)" {
			words[i] = CapitalizeEachWord(words[i])
			words = append(words[:i+1], words[i+2:]...)
			i--
		}

		if i >= 0 && i+1 < len(words) && words[i] == "(up," && strings.HasSuffix(words[i+1], ")") {
			nunberStr := words[i+1][:len(words[i+1])-1]
			number, err := strconv.Atoi(nunberStr)
			if number < 0 {
				log.Fatalf("Invalid negative number: (up, %v)", number)
			} else if number > i {
				log.Fatalf("over number : (up, %v)", number)
			}
			if err == nil {
				start := i - number
				for j := start; j <= i; j++ {
					if j >= 0 {
						words[j] = strings.ToUpper(words[j])
					}
				}
				words = append(words[:i], words[i+2:]...)
				if i != 0 {
					i -= 2
				}
			}
		}

		if i >= 0 && i+1 < len(words) && words[i] == "(low," && strings.HasSuffix(words[i+1], ")") {
			nunberStr := words[i+1][:len(words[i+1])-1]
			number, err := strconv.Atoi(nunberStr)
			if number < 0 {
				log.Fatalf("Invalid negative number: (low, %v)", number)
			} else if number > i {
				log.Fatalf("over number : (up, %v)", number)
			}
			if err == nil {
				start := i - number
				for j := start; j <= i; j++ {
					if j >= 0 {
						words[j] = strings.ToLower(words[j])
					}
				}
				words = append(words[:i], words[i+2:]...)
				if i != 0 {
					i -= 2
				}
			}
		}
		if i >= 0 && i+1 < len(words) && words[i] == "(cap," && strings.HasSuffix(words[i+1], ")") {
			nunberStr := words[i+1][:len(words[i+1])-1]
			number, err := strconv.Atoi(nunberStr)
			if number < 0 {
				log.Fatalf("Invalid negative number: (cap, %v)", number)
			} else if number > i {
				log.Fatalf("over number : (up, %v)", number)
			}
			if err == nil {
				start := i - number
				for j := start; j <= i; j++ {
					if j >= 0 {
						words[j] = CapitalizeEachWord(words[j])
					}
				}
				words = append(words[:i], words[i+2:]...)
				if i != 0 {
					i -= 2
				}
			}
		}

	}
	return strings.Join(words, " ")
}

func Vowel(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A" || words[i] == "'a" || words[i] == "'A") && len(words[i+1]) > 0 && isVowel(rune(words[i+1][0])) {
			if words[i] == "a" {
				words[i] = "an"
			} else if words[i] == "A" {
				words[i] = "An"
			} else if words[i] == "'a" {
				words[i] = "'an"
			} else if words[i] == "'A" {
				words[i] = "'An"
			}
		}
	}
	return strings.Join(words, " ")
}

func FixQuotes(Input string) string {
	var result []rune
	openQuote := false // it's
	for i, char := range Input {
		if i != 0 && i != len(Input)-1 && char == '\'' && unicode.IsLetter(rune(Input[i-1])) && unicode.IsLetter(rune(Input[i+1])) {
			result = append(result, '\'')
			continue
		}
		if i > 0 && char == ' ' && Input[i-1] == '\'' {   
			continue
		}
		if i+1 < len(Input) && char == ' ' && openQuote && Input[i+1] == '\'' {
			continue
		}

		if char == '\'' && !openQuote {
			result = append(result, '\'')
			openQuote = true
		} else if char == '\'' && openQuote {
			result = append(result, '\'')
			result = append(result, ' ')
			openQuote = false
		} else {
			result = append(result, (char))
		}
	}
	return string(result)
}

func FixPunc(Input string) string {
	var result []rune

	for i, char := range Input {
		if i < len(Input)-1 && char == ' ' && isPunctuation(rune(Input[i+1])) {
			continue
		}

		result = append(result, char)
		if i+1 < len(Input) && isPunctuation(char) && Input[i+1] != ' ' && !isPunctuation(rune(Input[i+1])) && Input[i+1] != '\'' {
			result = append(result, ' ')
		}
	}

	return string(result)
}
