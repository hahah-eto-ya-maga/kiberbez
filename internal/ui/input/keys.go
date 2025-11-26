package input

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadKeyForCipher(reader *bufio.Reader, cipherName string) any {
	switch strings.ToLower(cipherName) {
	case "caesar":
		for {
			fmt.Print("Введите ключ количество граней (целое число): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			k, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Неверный ввод, повторите попытку")
				continue
			}
			return k
		}
	}
	return nil
}
