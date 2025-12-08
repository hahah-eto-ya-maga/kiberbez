package cipher

import (
	"fmt"
	"kiberbez/internal/ui/colors"
	"strings"
)

type RC5Key struct {
	Key    []byte
	Rounds int
}

type RC5 struct {
	Key RC5Key
}

func NewRC5(key RC5Key) *RC5 {
	return &RC5{Key: key}
}

func (c *RC5) Name() string {
	return RC5Name
}

func (c *RC5) GetKey() string {
	var result []string
	result = append(result, fmt.Sprintf("Ключ%s:\n", colors.DEFAULT))
	result = append(result, fmt.Sprintf("%sКак строка%s: %s\n", colors.GREEN, colors.DEFAULT, c.Key.Key))
	result = append(result, fmt.Sprintf("%sКак байты (массив чисел)%s: %v\n", colors.GREEN, colors.DEFAULT, c.Key.Key))
	result = append(result, fmt.Sprintf("%sКак hex%s: %x\n", colors.GREEN, colors.DEFAULT, c.Key.Key))
	result = append(result, fmt.Sprintf("%sКоличество раундов%s: %d", colors.CYAN, colors.DEFAULT, c.Key.Rounds))

	return strings.Join(result, "")
}

func (c *RC5) SetKey(key any) {
	c.Key = key.(RC5Key)
}

func (c *RC5) Encrypt(text string) []string {
	var result []string

	return result
}

func (c *RC5) Decrypt(text string) []string {
	var result []string

	return result
}

func (c *RC5) Hack(text string) []string {
	var result []string
	result = append(result, "Тут ничего нет, а я думал должно быть, поэтому не стал переделывать архитектуру")

	return result
}
