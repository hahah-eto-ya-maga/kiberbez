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
	case cipher.RC5Name:
		for {
			fmt.Print("Введите ключ: ")
			input, _ := reader.ReadString('\n')
			key := strings.TrimSpace(input)

			var rounds int
			for {
				fmt.Print("Введите количество раундов (от 0 до 255): ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				var err error
				rounds, err = strconv.Atoi(input)
				if err != nil || rounds < 0 || rounds > 255 {
					fmt.Println("Неверный ввод, повторите попытку")
					continue
				}
				break
			}

			return cipher.RC5Key{
				Key:    []byte(key),
				Rounds: rounds,
			}
		}

	}
	return nil
}
