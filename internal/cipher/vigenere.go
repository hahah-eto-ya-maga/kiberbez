package cipher

import (
	"fmt"
	"strings"
)

type Vigenere struct {
	Key string
}

func NewVigenere(key string) *Vigenere {
	return &Vigenere{Key: key}
}

func (c *Vigenere) Name() string {
	return VigenereName
}

func (c *Vigenere) GetKey() string {
	return fmt.Sprintf("%s", c.Key)
}

func (c *Vigenere) SetKey(key any) {
	c.Key = key.(string)
}

func (c *Vigenere) Encrypt(text string) string {
	return encryptProcess(text, c.Key, false)
}

func (c *Vigenere) Decrypt(text string) string {
	return encryptProcess(text, c.Key, true)
}

func (c *Vigenere) Hack(text string) []string {
	var result []string

	return result
}

func buildKeyShifts(key string) []int {
	var result []int
	for _, r := range key {
		result = append(result, GetShift(r))
	}
	return result
}

func encryptProcess(text string, key string, invert bool) string {
	var result strings.Builder
	keyShifts := buildKeyShifts(key)
	keyIndex := 0

	for _, r := range text {
		shift := keyShifts[keyIndex]
		if invert {
			shift = -shift
		}

		result.WriteRune(ShiftRune(r, shift))

		keyIndex = (keyIndex + 1) % len(keyShifts)
	}

	return result.String()
}
