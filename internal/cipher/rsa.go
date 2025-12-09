package cipher

import (
	"fmt"
	"kiberbez/internal/ui/colors"
	"math/big"
	"strings"
)

type RSAKey struct {
	P *big.Int
	Q *big.Int
	E *big.Int
}

type RSA struct {
	Key RSAKey
}

func NewRSA(key RSAKey) *RSA {
	return &RSA{Key: key}
}

func (c *RSA) Name() string {
	return RC5Name
}

func (c *RSA) GetKey() string {
	var result []string
	result = append(result, fmt.Sprintf("Ключ%s:\n", colors.DEFAULT))
	result = append(result, fmt.Sprintf("%sp%s: %s\n", colors.GREEN, colors.DEFAULT, c.Key.P))
	result = append(result, fmt.Sprintf("%sq%s: %s\n", colors.GREEN, colors.DEFAULT, c.Key.Q))
	result = append(result, fmt.Sprintf("%se%s: %s\n", colors.GREEN, colors.DEFAULT, c.Key.E))

	return strings.Join(result, "")
}

func (c *RSA) SetKey(key any) {
	c.Key = key.(RSAKey)
}

func (c *RSA) Encrypt(text string) []string {
	var result []string

	return result
}

func (c *RSA) Decrypt(text string) []string {
	var result []string

	return result
}

func (c *RSA) Hack(text string) []string {
	var result []string
	result = append(result, "Тут ничего нет, а я думал должно быть, поэтому не стал переделывать архитектуру")

	return result
}
