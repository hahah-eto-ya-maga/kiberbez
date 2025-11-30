package input

import (
	"bufio"
	"fmt"
	"kiberbez/internal/cipher"
	"strconv"
	"strings"
)

func ReadKeyForCipher(reader *bufio.Reader, cipherName string) any {
	switch strings.ToLower(cipherName) {
	case cipher.CaesarName:
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
	case cipher.VigenereName:
		for {
			fmt.Print("Введите ключ слово (строка из БУКВ): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if !IsAlphaCyrillic(input) {
				fmt.Println("Неверный ввод, повторите попытку. Только кириллица")
				continue
			}

			return input
		}
	}
	return nil
}
