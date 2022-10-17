package hangman

import (
	"math/rand"
	"os"
	"time"
)

func SaveFileToTab() []string {
	var list []string
	var data []byte
	var tmp string
	data, _ = os.ReadFile("doc/words.txt")
	for _, letter := range data {
		if letter == '\n' {
			list = append(list, tmp)
			tmp = ""
		} else if letter != '\r' {
			tmp += string(letter)
		}
	}
	list = append(list, tmp)
	return list
}

func RandomWord(file []string) string {
	rand.Seed(time.Now().UnixNano())
	return file[rand.Intn(len(file)-1)]
}
