package main

import (
	"fmt"
	"hangman/hangman"
	"os"
)

func main() {
	var w hangman.Hang
	var letter string
	var tmp rune
	w.InitHang()
	w.GetRandomLetter()
	if len(os.Args) == 3 {
		if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {
			IsSave := w.GetData()
			if !IsSave {
				return
			}
		}
	} else if len(os.Args) != 1 {
		fmt.Println("invalid argument")
		return
	}
	fmt.Println("il vous reste", w.NbTry, "essais")
	fmt.Println(w.Guess)
	for w.Loop {
		fmt.Print("Choose a letter: ")
		fmt.Scanln(&letter)
		fmt.Println("\033[H\033[2J")
		if letter == "STOP" {
			w.SaveFile()
			break
		}
		if len(letter) == 1 {
			tmp = rune(letter[0])
			if tmp >= 97 && tmp <= 122 {
				if tmp >= 65 && tmp <= 90 {
					tmp += 32
					letter = ""
					letter += string(tmp)
				}
				w.CheckLetter(letter)
				w.CheckEnd()
			} else {
				fmt.Println("Not a letter")
			}
		} else if len(letter) == len(w.Word) {
			if letter == w.Word {
				for i := 0; i < len(w.Word); i++ {
					w.Guess[i] = string(w.Word[i])
				}
				fmt.Println("you win")
				w.Loop = false
			} else {
				if w.NbTry <= 2 {
					fmt.Println("You lose, the word was :", w.Word)
					w.Loop = false
				} else {
					w.I += 2
					w.NbTry -= 2
					fmt.Print("Wrong guess, ", w.NbTry, "/10 tries left\n")
				}
			}
		} else {
			fmt.Println("You must enter a letter or a guess")
		}
		fmt.Println(w.Guess)
		fmt.Println(w.José[w.I])
	}
}
