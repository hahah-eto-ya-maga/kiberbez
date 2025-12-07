package cipher

import (
	"fmt"
)

type RC5 struct {
	Key int
}

func NewRC5(shift int) *RC5 {
	return &RC5{Key: shift}
}

func (c *RC5) Name() string {
	return "RC5"
}

func (c *RC5) GetKey() string {
	return fmt.Sprintf("%d", c.Key)
}

func (c *RC5) SetKey(shift any) {
	c.Key = shift.(int)
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
	result = append(result, "Тут ничего нет, а я думал будет, поэтому не стал переделывать архитектуру")

	return result
}
