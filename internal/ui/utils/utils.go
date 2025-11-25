package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetUserChoice[T any](options []T, display func(T)) T {
	for i := range options {
		fmt.Printf("%d)\n", i+1)
		display(options[i])
		fmt.Println()
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nВыберите пример: ")
	input, _ := reader.ReadString('\n')
	input = strings.ReplaceAll(input, "\n", "")
	choice, _ := strconv.Atoi(input)

	ClearScreen()

	return options[choice-1]
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
