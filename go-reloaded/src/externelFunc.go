package goReloded

import (
	"strconv"
	"strings"
	"unicode"
)

func HexToDec(str string) string {
	s, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		return str
	}
	return strconv.FormatInt(s, 10)
}

func BinToDec(str string) string {
	s, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		return str
	}
	return strconv.FormatInt(s, 10)
}

func CapitalizeEachWord(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToLower(word)
			for j, r := range word {
				if unicode.IsLetter(r) {
					str := string(words[i][:j]) + strings.ToUpper(string(word[j])) + string(words[i][j+1:])
					words[i] = str
					break
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'H'
}

func ExtraSpaces(s string) string {
	var result strings.Builder
	inSpecial := false
	i := 0

	for i < len(s) {
		if s[i] == '(' {
			inSpecial = true
			result.WriteByte(s[i])
			i++
			for i < len(s) && s[i] == ' ' {
				result.WriteByte(s[i])
				i++
			}
			for i < len(s) && s[i] != ')' {
				result.WriteByte(s[i])
				i++
			}
			if i < len(s) && s[i] == ')' {
				result.WriteByte(s[i])
				i++
				inSpecial = false
			}
		} else if s[i] == ')' {
			inSpecial = false
			result.WriteByte(s[i])
			i++
		} else {
			if !inSpecial && unicode.IsSpace(rune(s[i])) {
				if result.Len() > 0 && result.String()[result.Len()-1] != ' ' {
					result.WriteByte(' ')
				}
			} else {
				result.WriteByte(s[i])
			}
			i++
		}
	}
	return strings.TrimRight(result.String(), " ")
}
func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func Domaine(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words)-1; i++ {
		if strings.HasSuffix(words[i], ".") && (words[i+1] == "com" || words[i+1] == "ma" || words[i+1] == "fr") {
			words[i] = words[i] + words[i+1]
			words = append(words[:i+1], words[i+2:]...)
		}
	}

	return strings.Join(words, " ")
}
