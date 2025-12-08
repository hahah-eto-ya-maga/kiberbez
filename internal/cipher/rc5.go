package cipher

import (
	"fmt"
	"kiberbez/internal/ui/colors"
	"strings"
)

const (
	w = 32
	P = 0xB7E15163 // P = Odd((e - 2) * 2^w)
	Q = 0x9E3779B9 // Q = Odd((φ - 1) * 2^w)
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
	result = append(result, fmt.Sprintf("%sКоличество раундов%s: %d\n", colors.CYAN, colors.DEFAULT, c.Key.Rounds))
	result = append(result, fmt.Sprintf("%sРазмер слова (w)%s: константно 32, для другого w не реализовано", colors.CYAN, colors.DEFAULT))

	return strings.Join(result, "")
}

func (c *RC5) SetKey(key any) {
	c.Key = key.(RC5Key)
}

func (c *RC5) Encrypt(text string) []string {
	var result []string

	key := c.Key.Key
	r := c.Key.Rounds
	t := 2 * (r + 1)

	S := make([]uint32, t)
	S[0] = P
	for i := 1; i < t; i++ {
		S[i] = S[i-1] + Q
	}

	L := makeWordArray(key)

	G, H := uint32(0), uint32(0)
	i, j := 0, 0
	n := max(3*len(S), 3*len(L))
	for k := 0; k < n; k++ {
		G = rotl32(S[i]+G+H, 3)
		S[i] = G
		H = rotl32(L[j]+G+H, G+H)
		L[j] = H
		i = (i + 1) % len(S)
		j = (j + 1) % len(L)
	}

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

func rotl32(x, y uint32) uint32 {
	return (x << (y & 31)) | (x >> (32 - (y & 31)))
}

func makeWordArray(key []byte) []uint32 {
	// w только 32
	cLen := (len(key) + 3) / 4
	L := make([]uint32, cLen)
	for i := 0; i < len(key); i++ {
		wordIndex := i / 4
		positionInWord := i % 4
		currentByte := uint32(key[i])
		L[wordIndex] = L[wordIndex] + (currentByte << (8 * positionInWord))
	}
	if len(L) == 0 {
		L = make([]uint32, 1)
		L[0] = 0
	}

	return L
}
