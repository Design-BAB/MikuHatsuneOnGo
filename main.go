package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

func play(note string, length int) {
	//fmt.Println("I am going to play the " + note + "note")
	var noteFeq float64
	if note == "ab" {
		noteFeq = 415.30
	} else if note == "bb" {
		noteFeq = 466.16
	} else if note == "c" {
		noteFeq = 523.25
	} else if note == "db" {
		noteFeq = 554.37
	} else if note == "eb" {
		noteFeq = 622.25
	} else if note == "f" {
		noteFeq = 698.46
	} else if note == "abHigh" {
		noteFeq = 830.61
	} else {
		noteFeq = 200
	}
	err := beeep.Beep(noteFeq, length)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(length/4) * time.Millisecond) // This pause comes after the first beep
}

func main() {
	//this is just for the enter "pause" function
	var input string
	// The first beep doesn't work for some reason so here we go
	//err := beeep.Beep(1000, 200)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//time.Sleep(600 * time.Millisecond) // This pause comes after the first beep
	//Now we can try and play something!
	//first bar
	play("bb", 500)
	play("f", 500)
	fmt.Println("Miku")
	play("bb", 500)
	play("f", 500)
	fmt.Println("Miku")
	//second bar
	play("eb", 250)
	play("f", 250)
	fmt.Print("You can ")
	play("eb", 250)
	play("db", 250)
	fmt.Print("call me... ")
	play("c", 500)
	play("ab", 500)
	fmt.Println("Miku!")
	//third bar
	play("bb", 500)
	fmt.Print("Blue ")
	play("f", 500)
	fmt.Println("hair")
	play("bb", 500)
	fmt.Print("Blue ")
	play("f", 500)
	fmt.Println("Tie")
	//forth bar
	play("eb", 250)
	play("f", 250)
	play("eb", 250)
	play("db", 250)
	play("abHigh", 500)
	play("f", 500)
	//next line starting on next to little 13
	play("bb", 500)
	play("f", 500)
	fmt.Print("Open ")
	play("bb", 500)
	play("f", 500)
	fmt.Println("secrete")
	//Second bar
	play("eb", 250)
	play("f", 250)
	play("eb", 250)
	fmt.Print("Anybody ")
	play("db", 250)
	fmt.Print("can ")
	play("c", 500)
	play("ab", 500)
	fmt.Println("find me!")
	//this is our pause before closing
	fmt.Println("Press Enter to continue...")
	fmt.Scanln(&input)
}
