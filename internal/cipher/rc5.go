package cipher

import (
	"encoding/binary"
	"fmt"
	"kiberbez/internal/ui/colors"
	"strconv"
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

	L := c.makeWordArray(key)

	G, H := uint32(0), uint32(0)
	i, j := 0, 0
	n := max(3*len(S), 3*len(L))
	for k := 0; k < n; k++ {
		G = c.rotl32(S[i]+G+H, 3)
		S[i] = G
		H = c.rotl32(L[j]+G+H, G+H)
		L[j] = H
		i = (i + 1) % len(S)
		j = (j + 1) % len(L)
	}

	data := []byte(text)

	if len(data)%8 != 0 {
		pad := 8 - (len(data) % 8)
		data = append(data, make([]byte, pad)...)
	}

	var ciphertext string
	for b := 0; b < len(data); b += 8 {
		A := binary.LittleEndian.Uint32(data[b : b+4])
		B := binary.LittleEndian.Uint32(data[b+4 : b+8])

		A += S[0]
		B += S[1]

		for round := 1; round <= r; round++ {
			B = c.rotl32(B^A, A) + S[2*round+1]
			A = c.rotl32(A^B, B) + S[2*round]
		}

		block := fmt.Sprintf("%08x%08x", A, B)
		ciphertext += block
	}
	result = append(result, ciphertext)

	return result
}

func (c *RC5) Decrypt(text string) []string {
	var result []string

	key := c.Key.Key
	r := c.Key.Rounds
	t := 2 * (r + 1)

	S := make([]uint32, t)
	S[0] = P
	for i := 1; i < t; i++ {
		S[i] = S[i-1] + Q
	}

	L := c.makeWordArray(key)

	G, H := uint32(0), uint32(0)
	i, j := 0, 0
	n := max(3*len(S), 3*len(L))
	for k := 0; k < n; k++ {
		G = c.rotl32(S[i]+G+H, 3)
		S[i] = G
		H = c.rotl32(L[j]+G+H, G+H)
		L[j] = H
		i = (i + 1) % len(S)
		j = (j + 1) % len(L)
	}

	var decoded []byte
	for b := 0; b < len(text); b += 16 {
		blockHex := text[b : b+16]

		A, _ := strconv.ParseUint(blockHex[:8], 16, 32)
		B, _ := strconv.ParseUint(blockHex[8:], 16, 32)

		A32 := uint32(A)
		B32 := uint32(B)

		for round := r; round >= 1; round-- {
			B32 = c.rotr32(B32-S[2*round+1], A32) ^ A32
			A32 = c.rotr32(A32-S[2*round], B32) ^ B32
		}

		B32 -= S[1]
		A32 -= S[0]

		buf := make([]byte, 8)
		binary.LittleEndian.PutUint32(buf[:4], A32)
		binary.LittleEndian.PutUint32(buf[4:], B32)

		decoded = append(decoded, buf...)
	}

	result = append(result, string(decoded))

	return result
}

func (c *RC5) Hack(text string) []string {
	var result []string
	result = append(result, "Тут ничего нет, а я думал должно быть, поэтому не стал переделывать архитектуру")

	return result
}

func (c *RC5) rotl32(x, y uint32) uint32 {
	return (x << (y & 31)) | (x >> (32 - (y & 31)))
}

func (c *RC5) rotr32(x, y uint32) uint32 {
	return (x >> (y & 31)) | (x << (32 - (y & 31)))
}

func (c *RC5) makeWordArray(key []byte) []uint32 {
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
