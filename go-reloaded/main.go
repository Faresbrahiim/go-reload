package main

import (
	"bufio"
	goReloded "goReloded/src"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provide the correct number of arguments")
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]
	if !strings.HasSuffix(inputFileName, ".txt") || !strings.HasSuffix(outputFileName, ".txt") {
		log.Fatal("Check your arguments, make sure they end with .txt")
	}

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal("error in oppening:", err)
	}
	defer file.Close()

	output, err := os.Create(outputFileName)
	if err != nil {
		log.Fatal("errr in creating", err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(output)

	for scanner.Scan() {
		line := scanner.Text()
		stringConverted := goReloded.StringConvert(line)
		fixPQuots := goReloded.FixQuotes(stringConverted)
		fixpunc := goReloded.FixPunc(fixPQuots)
		FixVowel := goReloded.Vowel(fixpunc)
		finalTxt := goReloded.Domaine(FixVowel)

		_, err := writer.WriteString(finalTxt + "\n")
		if err != nil {
			log.Fatal("Error writing :", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading", err)
	}

	err = writer.Flush()  // forced any data to be written if we don't use  may some data  don' t bw written
	if err != nil {
		log.Fatal("Error flushing the writer:", err)
	}	
}
