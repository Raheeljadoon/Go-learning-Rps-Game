package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var p1Choice int
	var p2Choice int
	var choice string
	score := map[string]int{"player1": 0, "computer": 0}
	for {
		fmt.Println("1 : Rock\n2 : Paper\n3 : Scissor")

		fmt.Println("Game Begin...\nPlayer1 Choose from Above three")
		_, _ = fmt.Scan(&p1Choice)
		rand.Seed(time.Now().Unix())
		p2Choice = rand.Intn(3) + 1

		fmt.Println("Player 1 choose -> ", p1Choice, "\nComputer 2 choose ->", p2Choice)
		if p1Choice == 1 && p2Choice == 3 || p1Choice == 2 && p2Choice == 1 || p1Choice == 3 && p2Choice == 2 {
			fmt.Println("Player 1 Wins")
			score["player1"] += 1

		} else if p1Choice == p2Choice {
			fmt.Println("Its a Tie")
		} else {
			fmt.Println("Computer Wins")
			score["computer"] += 1
		}
		fmt.Println("Do you want to play again : y or n !")
		_, _ = fmt.Scan(&choice)
		if choice == "y" {
			continue
		} else {
			break
		}
	}
	fmt.Println("Game End Thanks\nScore was ->", "computer :", score["computer"], "player1 :", score["player1"])
}
