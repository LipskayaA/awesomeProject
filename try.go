// try to create my first simple game

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func reader() int {
	fmt.Print("Enter a grade: ")
	reader1 := bufio.NewReader(os.Stdin)
	input, err := reader1.ReadString('\n')

	if err != nil {
		return 0
	}
	input = strings.TrimSpace(input)
	number, err1 := strconv.Atoi(input)
	if err1 != nil {
		return 0

	} else {
		return number
	}

}

func main() {
	fmt.Println("I made a number between 1 and 50. Let`s try to guess!1")
	myNumber := rand.Intn(50) + 1
	for x := 1; true; x++ {
		input := reader()
		if input > myNumber {
			fmt.Println("It`s too High, try again!")
		}
		if input < myNumber {
			fmt.Println("It`s too Low, try again")
		}
		if input == myNumber {
			fmt.Println("You`re right!")
			break
		}
	}

}
