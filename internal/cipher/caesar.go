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

func (c *Caesar) Encrypt(text string) []string {
	var result []string
	result = append(result, shiftedEncrypt(text, c.Shift))
	return result
}

func (c *Caesar) Decrypt(text string) []string {
	var result []string
	result = append(result, shiftedEncrypt(text, -c.Shift))
	return result
}

func (c *Caesar) Hack(text string) []string {
	var result []string
	result = append(result, "["+colors.CYAN+"Метод перебора сдвигов"+colors.DEFAULT+"]")

	for shift := 1; shift < 34; shift++ {
		temp := shiftedEncrypt(text, -shift)
		result = append(result, fmt.Sprintf("%s [сдвиг %d]", temp, shift))
	}
	result = append(result, colors.CYAN+"В результате взлома правильный текст находится там, где он читается осмысленно"+colors.DEFAULT)

	return result
}

func shiftedEncrypt(text string, shift int) string {
	var result strings.Builder
	for _, r := range text {
		result.WriteRune(ShiftRune(r, shift))
	}
	return result.String()
}
