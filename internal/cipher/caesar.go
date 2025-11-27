package cipher

import (
	"fmt"
	"kiberbez/internal/ui/colors"
	"strings"
)

type Caesar struct {
	Shift int
}

func NewCaesar(shift int) *Caesar {
	return &Caesar{Shift: shift}
}

func (c *Caesar) Name() string {
	return "Caesar"
}

func (c *Caesar) GetKey() string {
	return fmt.Sprintf("%d", c.Shift)
}

func (c *Caesar) SetKey(shift any) {
	c.Shift = shift.(int)
}

func (c *Caesar) Encrypt(text string) string {
	var result strings.Builder

	for _, char := range text {
		switch {
		case 'а' <= char && char <= 'я':
			shift := (c.Shift%32 + 32) % 32
			newChar := rune((int(char-'а')+shift)%32 + int('а'))
			result.WriteRune(newChar)
		case 'А' <= char && char <= 'Я':
			shift := (c.Shift%32 + 32) % 32
			newChar := rune((int(char-'А')+shift)%32 + int('А'))
			result.WriteRune(newChar)
		case 'a' <= char && char <= 'z':
			shift := (c.Shift%27 + 27) % 27
			newChar := rune((int(char-'a')+shift)%27 + int('a'))
			result.WriteRune(newChar)
		case 'A' <= char && char <= 'Z':
			shift := (c.Shift%27 + 27) % 27
			newChar := rune((int(char-'A')+shift)%27 + int('A'))
			result.WriteRune(newChar)
		default:
			result.WriteRune(char)
		}
	}
	return result.String()
}

func (c *Caesar) Decrypt(text string) string {
	return (&Caesar{Shift: -c.Shift}).Encrypt(text)
}

func (c *Caesar) Hack(text string) []string {
	var result []string
	result = append(result, "["+colors.CYAN+"Метод перебора сдвигов"+colors.DEFAULT+"]")

	alphabetSize := detectAlphabet(text)

	for shift := 0; shift < alphabetSize; shift++ {
		temp := Caesar{Shift: -shift} // ну это плохо прям конечно
		result = append(result, fmt.Sprintf("%s [сдвиг %d]", temp.Encrypt(text), shift))
	}

	result = append(result, colors.YELLOW+"В результате взлома правильный текст находится там, где он читается осмысленно"+colors.DEFAULT)

	return result
}

func detectAlphabet(text string) int {
	for _, ch := range text {
		switch {
		case 'а' <= ch && ch <= 'я', 'А' <= ch && ch <= 'Я':
			return 32
		case 'a' <= ch && ch <= 'z', 'A' <= ch && ch <= 'Z':
			return 27
		}
	}
	return 32
}
