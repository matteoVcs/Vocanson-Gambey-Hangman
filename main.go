package main

import (
	"fmt"
	"hangman/hangman"
)

func main() {
	var w hangman.Hang
	var letter string
	var tmp rune
	w.InitHang()
	w.GetRandomLetter()
	for w.Loop {
		fmt.Print("Choose a letter: ")
		fmt.Scanln(&letter)
		fmt.Println("\033[H\033[2J")
		if len(letter) == 1 {
			tmp = rune(letter[0])
			if tmp >= 65 && tmp <= 90 || tmp >= 97 && tmp <= 122 {
			} else {
				fmt.Println("Not a letter")
			}
		} else {
			fmt.Println("Too long")
		}
		w.CheckLetter(letter)
		w.CheckEnd()
		fmt.Println(w.Guess)
		fmt.Println(w.José[w.I])
	}
}
