package menuActions

import (
	"bufio"
	"fmt"
	"kiberbez/internal/cipher"
	"kiberbez/internal/ui/colors"
	"kiberbez/internal/ui/input"
	"strings"
)

func EncryptCase(reader *bufio.Reader, c cipher.Cipher) {
	fmt.Print("Введите текст для зашифрования: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(colors.CYAN+"Результат зашифрования: "+colors.DEFAULT, c.Encrypt(text))
}

func DecryptCase(reader *bufio.Reader, c cipher.Cipher) {
	fmt.Print("Введите текст для дешифрования: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(colors.CYAN+"Результат дешифрования: "+colors.DEFAULT, c.Decrypt(text))
}

func HackCase(reader *bufio.Reader, c cipher.Cipher) {
	fmt.Print("Введите текст для взлома: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(colors.CYAN + "Результат взлома: " + colors.DEFAULT)
	for _, r := range c.Hack(text) {
		fmt.Println(" -", r)
	}
}

func ReadKeyForCipherCase(reader *bufio.Reader, c cipher.Cipher) {
	newKey := input.ReadKeyForCipher(reader, c.Name())
	c.SetKey(newKey)
}
