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

	for _, r := range text {
		result.WriteRune(shiftRune(r, c.Shift))
	}
	return result.String()
}

func (c *Caesar) Decrypt(text string) string {
	return (&Caesar{Shift: -c.Shift}).Encrypt(text)
}

func (c *Caesar) Hack(text string) []string {
	var result []string
	result = append(result, "["+colors.CYAN+"Метод перебора сдвигов"+colors.DEFAULT+"]")

	for shift := 1; shift < 34; shift++ {
		temp := (&Caesar{Shift: -shift}).Encrypt(text) // ну это плохо прям конечно
		result = append(result, fmt.Sprintf("%s [сдвиг %d]", temp, shift))
	}
	result = append(result, colors.CYAN+"В результате взлома правильный текст находится там, где он читается осмысленно"+colors.DEFAULT)

	return result
}

func shiftRune(r rune, shift int) rune {
	switch {
	case 'а' <= r && r <= 'я':
		return 'а' + rune((int(r-'а')+shift+32)%32)
	case 'А' <= r && r <= 'Я':
		return 'А' + rune((int(r-'А')+shift+32)%32)
	}

	switch {
	case 'a' <= r && r <= 'z':
		return 'a' + rune((int(r-'a')+shift+26)%26)
	case 'A' <= r && r <= 'Z':
		return 'A' + rune((int(r-'A')+shift+26)%26)
	}

	return r
}
