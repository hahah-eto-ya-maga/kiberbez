package cipher

import "fmt"

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
	return text
}

func (c *Caesar) Decrypt(text string) string {
	return text
}

func (c *Caesar) Hack(text string) []string {
	return []string{"[caesar bruteforce result]"}
}
