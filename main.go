package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ввод
func input() (vy, vx int) {
	for {
		var integer string
		in := bufio.NewReader(os.Stdin)
		integer, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
		}
		newInteger := strings.TrimSpace(integer)
		stringSlice := strings.Split(newInteger, " ")
		if len(stringSlice) > 2 {
			fmt.Println("Неверно,введите от 1 до 3: ")
			continue
		}
		vy, _ = strconv.Atoi(stringSlice[0])
		vx, _ = strconv.Atoi(stringSlice[1])
		if vy > 3 || vx > 3 || vy < 1 || vx < 1 {
			fmt.Println("Неверно,введите от 1 до 3: ")
			continue
		}
		return vy, vx
	}
}

// Поле для вывода
func field(copySlice [][]string) {
	for y := 0; y < len(copySlice); y++ {
		for x := 0; x < len(copySlice); x++ {
			fmt.Print(copySlice[y][x])
		}
		fmt.Print("\n")
	}
}

func switchPlayer(p1 string) string {
	if p1 == "X" {
		return "O"
	}
	return "X"
}

// Сама игра
func game() string {
	p1 := "X"
	slice := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Print("Вводите числа сначала y (1-3),после x (1-3)\n")
	freeCeils := 9
game:
	for freeCeils > 0 {
		fmt.Printf("Ходит %v\n", p1)
		v, q := input()
		v -= 1
		q -= 1
		if slice[v][q] == "X" || slice[v][q] == "O" {
			fmt.Print("Клетка занята,введите заново y и x (1-3): \n")
			field(slice)
			continue game
		}
		slice[v][q] = p1
		field(slice)
		if checkWinner(slice) != "" {
			fmt.Println("Победитель: " + p1)
			break game
		}
		p1 = switchPlayer(p1)
		freeCeils--
	}
	if freeCeils == 0 {
		return "Ничья!"
	} else {
		return ""
	}
}

// Проверка победителя
func checkWinner(checkfield [][]string) string {
	for i := 0; i < 3; i++ {
		if checkfield[i][0] != "_" && checkfield[i][0] == checkfield[i][1] && checkfield[i][1] == checkfield[i][2] {
			return "Победитель:"
		}
	}
	for j := 0; j < 3; j++ {
		if checkfield[0][j] != "_" && checkfield[0][j] == checkfield[1][j] && checkfield[1][j] == checkfield[2][j] {
			return "Победитель:"
		}
	}
	if checkfield[0][0] != "_" && checkfield[0][0] == checkfield[1][1] && checkfield[1][1] == checkfield[2][2] {
		return "Победитель:"
	}
	if checkfield[0][2] != "_" && checkfield[0][2] == checkfield[1][1] && checkfield[1][1] == checkfield[2][0] {
		return "Победитель:"
	}
	return ""
}

func main() {
	for {
		fmt.Println(game())
	}
}
