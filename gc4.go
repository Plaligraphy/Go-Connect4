package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ln1 [6]int
var ln2 [6]int
var ln3 [6]int
var ln4 [6]int
var ln5 [6]int
var curPlayer int

func main() {
	var in string
	fmt.Println("Connect Four (Terminal)")
	fmt.Println("s. Start Game")
	fmt.Println("ex. Exit")
	fmt.Scanln(&in)
	switch in {
	case "s":
		run()
	case "ex":
		os.Exit(0)
	}
}
func run() {
	var isRunning bool

	isRunning = true
	if !isRunning {
		os.Exit(0)
	} else {
		//for isRunning == true {
		draw()
		logic()
	}
}
func draw() {
	cls()
	fmt.Println(curPlayer)
	in1 := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(in1)
	fmt.Println(ln1)
	fmt.Println(ln2)
	fmt.Println(ln3)
	fmt.Println(ln4)
	fmt.Println(ln5)
}
func logic() {
	if curPlayer == 0 {
		curPlayer = 1
	}
	if curPlayer == 1 || curPlayer == 2 {
		restart()
	}
	var rowSelect int
	if curPlayer == 1 {
		fmt.Println("Player 1 Turn")
		fmt.Println("Row?: (0-5) 25 to exit")
		fmt.Scanln(&rowSelect)
	}
	if curPlayer == 2 {
		fmt.Println("Player 2 Turn")
		fmt.Println("Row?: (0-5) 25 to exit")
		fmt.Scanln(&rowSelect)
	}

	if curPlayer == 3 {
		rand.Seed(time.Now().UnixNano())
		max := 3
		min := 0
		rand.Seed(time.Now().UnixNano())
		val := rand.Intn(max-min) + min
		col(val)
		draw()
		//Implement AI here
	}
	switch rowSelect {
	case 42:
		clearBoard()
		draw()
		logic()
	case 25:
		os.Exit(0)
	case 0:
		//USE A SELECTOR FOR WHICH LINE
		col(0)
		draw()
		logic()
	case 1:
		col(1)
		draw()
		logic()
	case 2:
		col(2)
		draw()
		logic()
	case 3:
		col(3)
		draw()
		logic()
	case 4:
		col(4)
		draw()
		logic()
	case 5:
		col(5)
		draw()
		logic()
	}
}
func col(row int) {
	//Works idk how but hey works
	var cursorX int
	var cursorPos int
	if ln5[row] == 1 || ln5[row] == 2 {
		cursorX = 0
		cursorX++
		if ln4[row] == 1 || ln4[row] == 2 {
			cursorX++
			if ln3[row] == 1 || ln3[row] == 2 {
				cursorX++
				if ln2[row] == 1 || ln2[row] == 2 {
					cursorX++
					if ln1[row] == 1 || ln1[row] == 2 {
						checkWin()
						fmt.Println("Invalid!, Try a different row!")
						logic()
					} else {
						cursorPos = 1
					}
				} else {
					cursorPos = 2
				}
			} else {
				cursorPos = 3
			}

		} else {
			cursorPos = 4
		}

	} else {
		cursorPos = 5
	}
	switch cursorPos {
	case 1:
		ln1[row] = curPlayer
	case 2:
		ln2[row] = curPlayer
	case 3:
		ln3[row] = curPlayer
	case 4:
		ln4[row] = curPlayer
	case 5:
		ln5[row] = curPlayer
	}
	if cursorX >= 5 {

	}
}
func cls() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
func clearBoard() {
	for i := 0; i <= 5; i++ {
		ln1[i] = 0
		ln2[i] = 0
		ln3[i] = 0
		ln4[i] = 0
		ln5[i] = 0
	}

}
func restart() {
	if curPlayer == 1 {
		curPlayer = 2
	} else {
		curPlayer = 1
	}
}
func checkWin() {
	//Lose / Win conditions
	if ln1[0] == 1 && ln1[1] == 1 && ln1[2] == 1 && ln1[3] == 1 && ln1[4] == 1 && ln1[5] == 1 {
		win()
	}
	if ln2[0] == 1 && ln2[1] == 1 && ln2[2] == 1 && ln2[3] == 1 && ln2[4] == 1 && ln2[5] == 1 {
		win()
	}
	if ln3[0] == 1 && ln3[1] == 1 && ln3[2] == 1 && ln3[3] == 1 && ln3[4] == 1 && ln3[5] == 1 {
		win()
	}
	if ln4[0] == 1 && ln4[1] == 1 && ln4[2] == 1 && ln4[3] == 1 && ln4[4] == 1 && ln4[5] == 1 {
		win()
	}
	if ln5[0] == 1 && ln5[1] == 1 && ln5[2] == 1 && ln5[3] == 1 && ln5[4] == 1 && ln5[5] == 1 {
		win()
	}
	if ln1[0] == 1 && ln2[0] == 1 && ln3[0] == 1 && ln4[0] == 1 && ln5[0] == 1 || ln1[0] == 2 && ln2[0] == 2 && ln3[0] == 2 && ln4[0] == 2 && ln5[0] == 2 {
		win()

	}
}
func win() {
	fmt.Println("You Win!")
	time.Sleep(2 * time.Second)
	os.Exit(0)
}

//Need to add diagonal win detection
