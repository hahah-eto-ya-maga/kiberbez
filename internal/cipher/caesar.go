package cipher

type Caesar struct {
	Shift int
}

func NewCaesar(shift int) *Caesar {
	return &Caesar{Shift: shift}
}

func (c *Caesar) Name() string {
	return "Caesar"
}

func (c *Caesar) Encrypt(text string) string {
	return "[caesar encrypted: " + text + "]"
}

func (c *Caesar) Decrypt(text string) string {
	return "[caesar decrypted: " + text + "]"
}

func (c *Caesar) Bruteforce(text string) []string {
	return []string{"[caesar bruteforce result]"}
}
