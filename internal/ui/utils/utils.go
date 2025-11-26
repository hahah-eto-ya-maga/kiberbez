package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserChoice[T any](options []T, display func(T)) T {
	for {
		for i := range options {
			fmt.Printf("%d) ", i+1)
			display(options[i])
			fmt.Println()
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nВыберите: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice <= 0 || choice > len(options) {
			fmt.Println("Неверный ввод, повторите попытку")
			continue
		}

		return options[choice-1]
	}

}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
