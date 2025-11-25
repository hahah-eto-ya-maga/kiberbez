package cipher

type Cipher interface {
	Encrypt(string) string
	Decrypt(string) string
	Bruteforce(string) []string
	Name() string
}

func NewCipher(name string, key any) Cipher {
	switch name {
	case "caesar":
		if k, ok := key.(int); ok {
			return NewCaesar(k)
		}
	}
	return nil
}

func ListCiphers() []string {
	return []string{"caesar"}
}
