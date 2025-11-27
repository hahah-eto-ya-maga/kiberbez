package main

import (
	"fmt"
	"strings"
)

func encrypt(plaintext string, shift int) string {
	var result strings.Builder
	shift = (shift%32 + 32) % 32

	for _, char := range plaintext {
		switch {
		case 'а' <= char && char <= 'я':
			newChar := rune((int(char-'а')+shift)%32 + int('а'))
			result.WriteRune(newChar)
		case 'А' <= char && char <= 'Я':
			newChar := rune((int(char-'А')+shift)%32 + int('А'))
			result.WriteRune(newChar)
		default:
			result.WriteRune(char)
		}
	}
	return result.String()
}

func decrypt(ciphertext string, shift int) string {
	return encrypt(ciphertext, 32-shift)
}

func crack(ciphertext string) map[int]string {
	possiblePlaintexts := make(map[int]string)
	for s := 1; s <= 31; s++ {
		plaintext := decrypt(ciphertext, s)
		possiblePlaintexts[s] = plaintext
	}
	return possiblePlaintexts
}

func main() {
	originalText := "Привет вам, Алексей Сергеевич."
	shiftKey := 13

	encryptedText := encrypt(originalText, shiftKey)
	fmt.Printf("--- Шифрование ---\n")
	fmt.Printf("Исходный текст: %s\n", originalText)
	fmt.Printf("Сдвиг (K): %d\n", shiftKey)
	fmt.Printf("Зашифрованный: %s\n\n", encryptedText)

	decryptedText := decrypt(encryptedText, shiftKey)
	fmt.Printf("--- Дешифрование ---\n")
	fmt.Printf("Зашифрованный: %s\n", encryptedText)
	fmt.Printf("Сдвиг (K): %d\n", shiftKey)
	fmt.Printf("Расшифрованный: %s\n\n", decryptedText)

	fmt.Printf("--- Брутфорс ---\n")
	fmt.Printf("Попытка взлома текста: %s\n", encryptedText)

	crackedOptions := crack(encryptedText)
	for shift, plaintext := range crackedOptions {
		if shift == shiftKey {
			fmt.Printf("-> Сдвиг K=%-2d: %s\n", shift, plaintext)
		} else {
			fmt.Printf("Сдвиг K=%-2d: %s\n", shift, plaintext)
		}
	}

	fmt.Println("\nВ результате взлома правильный текст находится там, где он читается осмысленно.")
}
