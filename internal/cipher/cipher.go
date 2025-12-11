package cipher

const (
	CaesarName   string = "caesar"
	VigenereName string = "vigenere"
	RC5Name      string = "rc5"
	RSAName      string = "rsa"
)

type Cipher interface {
	Encrypt(string) []string
	Decrypt(string) []string
	Hack(string) []string
	Name() string
	GetKey() string
	SetKey(any)
}

func NewCipher(name string, key any) Cipher {
	switch name {
	case CaesarName:
		if k, ok := key.(int); ok {
			return NewCaesar(k)
		}
	case VigenereName:
		if k, ok := key.(string); ok {
			return NewVigenere(k)
		}
	case RC5Name:
		if k, ok := key.(RC5Key); ok {
			return NewRC5(k)
		}
	case RSAName:
		if k, ok := key.(RSAProps); ok {
			return NewRSA(k)
		}
	}
	return nil
}

func ListCiphers() []string {
	return []string{CaesarName, VigenereName, RC5Name, RSAName}
}
