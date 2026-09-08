package cipher

import (
	"fmt"
	"kiberbez/internal/ui/colors"
	"math/big"
	"strings"
)

type publicKey struct {
	e *big.Int
	n *big.Int
}

type privateKey struct {
	d *big.Int
	n *big.Int
}

type RSAKey struct {
	prKey  privateKey
	pubKey publicKey
}

type RSA struct {
	Key RSAKey
}

type RSAProps struct {
	P *big.Int
	Q *big.Int
	E *big.Int
}

func NewRSA(key RSAProps) *RSA {
	n := new(big.Int).Mul(
		key.P, key.Q,
	)
	pMinus1 := new(big.Int).Sub(key.P, big.NewInt(1))
	qMinus1 := new(big.Int).Sub(key.Q, big.NewInt(1))
	phi := new(big.Int).Mul(pMinus1, qMinus1)

	d := new(big.Int).ModInverse(key.E, phi)

	return &RSA{Key: RSAKey{
		prKey: privateKey{
			d: d,
			n: n,
		},
		pubKey: publicKey{
			e: key.E,
			n: n,
		},
	}}
}

func (c *RSA) Name() string {
	return RSAName
}

func (c *RSA) GetKey() string {
	var result []string
	result = append(result, fmt.Sprintf("Публичный ключ%s: (%v, %v)\n", colors.DEFAULT, c.Key.pubKey.e, c.Key.pubKey.n))
	result = append(result, fmt.Sprintf("%sПриватный ключ%s: (%v, %v)\n", colors.GREEN, colors.DEFAULT, c.Key.prKey.d, c.Key.prKey.n))

	return strings.Join(result, "")
}

func (c *RSA) SetKey(key any) {
	c.Key = key.(RSAKey)
}

func (c *RSA) Encrypt(text string) []string {
	var result []string

	e := c.Key.prKey.d
	n := c.Key.pubKey.n

	blockSize := (n.BitLen() - 1) / 8
	if blockSize < 1 {
		blockSize = 1
	}

	textBytes := []byte(text)

	var encryptedBlocks []string
	for i := 0; i < len(textBytes); i += blockSize {
		end := i + blockSize
		if end > len(textBytes) {
			end = len(textBytes)
		}

		chunk := textBytes[i:end]

		M := new(big.Int).SetBytes(chunk)

		S := new(big.Int).Exp(M, e, n)

		encryptedBlocks = append(encryptedBlocks, S.String())
	}

	result = append(result, strings.Join(encryptedBlocks, ":"))

	return result
}

func (c *RSA) Decrypt(text string) []string {
	var result []string

	d := c.Key.pubKey.e
	n := c.Key.pubKey.n

	if text == "" {
		return []string{""}
	}

	blocks := strings.Split(text, ":")

	var decryptedBytes []byte
	for _, blockStr := range blocks {
		S := new(big.Int)
		S.SetString(blockStr, 10)

		M := new(big.Int).Exp(S, d, n)

		decryptedBytes = append(decryptedBytes, M.Bytes()...)
	}

	result = append(result, string(decryptedBytes))

	return result
}

func (c *RSA) Hack(text string) []string {
	var result []string
	result = append(result, "Тут ничего нет, а я думал должно быть, поэтому не стал переделывать архитектуру")

	return result
}
