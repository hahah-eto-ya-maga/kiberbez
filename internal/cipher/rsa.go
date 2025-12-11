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
	return RC5Name
}

func (c *RSA) GetKey() string {
	var result []string
	result = append(result, fmt.Sprintf("Публичный ключ%s: (%v, %v)\n", colors.DEFAULT, c.Key.pubKey.e, c.Key.pubKey.n))
	result = append(result, fmt.Sprintf("Приватный ключ%s: (%v, %v)\n", colors.DEFAULT, c.Key.prKey.d, c.Key.prKey.n))

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
