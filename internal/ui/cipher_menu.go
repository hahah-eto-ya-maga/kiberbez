package ui

import (
	"bufio"
	"fmt"
	"kiberbez/internal/cipher"
	"kiberbez/internal/ui/input"
	"kiberbez/internal/ui/utils"
	"os"
)

func SelectCipher() cipher.Cipher {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("Доступные шифры:")
		cipherName := utils.GetUserChoice(cipher.ListCiphers(), func(name string) {
			fmt.Print(name)
		})

		key := input.ReadKeyForCipher(reader, cipherName)

		c := cipher.NewCipher(cipherName, key)
		if c == nil {
			fmt.Println("Неизвестный шифр")
			continue
		}

		return c
	}
}
