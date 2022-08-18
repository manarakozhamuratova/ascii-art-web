package service

import (
	"bytes"
	"os"
	"strings"
)

const (
	fontsPath     = "static/fonts/"
	fonstStandard = "standard.txt"
	fileSuffix    = "txt"
	countLines    = 856
)

func fileRead(fileName string) (string, error) {
	data, err := os.ReadFile(fontsPath + fileName)
	if err != nil {
		return "", err
	}
	str1 := bytes.Split(data, []byte("\n"))
	if len(str1) != countLines {
		return "", err
	}
	return string(data), nil
}

func distributionByAscii(str, data string) (map[rune]string, []string) {
	mapa := make(map[rune]string)
	temp := ""
	count := 0
	keys := rune(32)
	strand := strings.ReplaceAll(str, "\r\n", "\n")
	splitText := strings.Split(strand, "\n")

	for _, words := range data {
		temp += string(words)
		if words == '\n' {
			count++
		}
		if count == 9 {
			mapa[keys] = temp[1 : len(temp)-1]
			temp = ""
			count = 0
			keys++
		}

	}
	return mapa, splitText
}

func compareWithTheMessage(str string, letters map[rune]string) string {
	var template [8]string
	if len(str) == 0 {
		return "\n"
	}
	for _, w := range str {
		for index, value := range strings.Split(letters[w], "\n") {
			template[index] += value
		}
	}
	result := ""
	for _, words := range template {
		result += string(words) + "\n"
	}

	return result
}

func returnTheAnswer(str []string, mapa map[rune]string) string {
	result := ""
	for _, words := range str {
		result += compareWithTheMessage(words, mapa)
	}
	return result
}

func getFontName(banner string) string {
	return banner + ".txt"
}

func App(message, banner string) (string, error) {
	nameFile := getFontName(banner)
	data, err := fileRead(nameFile)
	if err != nil {
		return "", err
	}
	mapa, splitText := distributionByAscii(message, data)
	res := returnTheAnswer(splitText, mapa)
	return res, nil
}
