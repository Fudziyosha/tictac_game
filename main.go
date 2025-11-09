package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Ввод
func input() (vy, vx int) {
	for {
		vy = checkInputOne(vy)
		vx = checkInputTwo(vx)
		if vy > 3 || vx > 3 || vy < 1 || vx < 1 {
			fmt.Println("Неверно,введите от 1 до 3: ")
			continue
		}
		return vy, vx
	}
}

// Проверка ввода и конвертация в int
func checkInputOne(num int) int {
	bn := bufio.NewScanner(os.Stdin)
	if bn.Scan() {
		line := bn.Text()
		num, _ = strconv.Atoi(line)
	}
	return num
}

func checkInputTwo(num2 int) int {
	bn := bufio.NewScanner(os.Stdin)
	if bn.Scan() {
		line := bn.Text()
		num2, _ = strconv.Atoi(line)
	}
	return num2
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

// Сама игра
func game() {
loop:
	for {
		slice := [][]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		}
		fmt.Print("Вводите числа сначала y (1-3),после x (1-3)\n")
		p1 := "X"
		freeCeils := 9
	game:
		for freeCeils > 0 {
			fmt.Printf("Ходит %v\n", p1)
			v, q := input()
			v -= 1
			q -= 1
			for y := 0; y < len(slice); y++ {
				for x := 0; x < len(slice); x++ {
					if y == v && x == q {
						if slice[y][x] == "X" || slice[y][x] == "O" {
							fmt.Print("Клетка занята,введите заново y и x (1-3): \n")
							field(slice)
							continue game
						}
						slice[y][x] = p1
						field(slice)
					}
				}
				if checkWinner(slice) != "" {
					fmt.Println("Победитель: " + p1)
					continue loop
				}
			}
			if p1 == "X" {
				p1 = "O"
			} else {
				p1 = "X"
			}
			freeCeils--
			if freeCeils < 1 {
				fmt.Println("Ничья!")
				continue loop
			}
		}
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
	game()
}
