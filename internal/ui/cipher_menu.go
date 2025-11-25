package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"kiberbez/internal/cipher"
)

func SelectCipher() cipher.Cipher {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Available ciphers:")
		for _, name := range cipher.ListCiphers() {
			fmt.Println(" -", name)
		}

		fmt.Print("Choose cipher: ")
		cipherName, _ := reader.ReadString('\n')
		cipherName = strings.TrimSpace(cipherName)

		fmt.Print("Enter integer key: ")
		keyStr, _ := reader.ReadString('\n')
		keyStr = strings.TrimSpace(keyStr)

		key, err := strconv.Atoi(keyStr)
		if err != nil {
			fmt.Println("Key must be a number")
			continue
		}

		c := cipher.NewCipher(cipherName, key)
		if c == nil {
			fmt.Println("Unknown cipher")
			continue
		}

		return c
	}
}
