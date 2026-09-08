package cipher

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

type StreamRC5 struct {
	// вообще кстати можно сделать универсальную потоковую строктуру. Только непонятно как ключи брать пока. Может и просто всё на свежий взгляд будет
	rc5     *RC5
	nonce   uint32
	counter uint32
}

func NewStreamRC5(rc5 *RC5) *StreamRC5 {
	return &StreamRC5{
		rc5:     rc5,
		nonce:   uint32(time.Now().UnixNano() & 0xffffffff),
		counter: 0,
	}
}

func (c *StreamRC5) Name() string {
	return StreamRC5Name
}

func (c *StreamRC5) GetKey() string {
	return c.rc5.GetKey()
}

func (c *StreamRC5) SetKey(key any) {
	c.rc5.Key = key.(RC5Key)
}

func (c *StreamRC5) Encrypt(text string) []string {
	var result []string
	result = append(result, fmt.Sprintf("nonce=%08x", c.nonce))

	data := []byte(text)
	var ks []byte
	c.counter = 0

	out := make([]byte, len(data))
	for i := range data {
		if i%8 == 0 { // w = 32 поэтому так
			ks = c.keystreamBlock()
		}
		out[i] = data[i] ^ ks[i%8]
	}

	result = append(result, fmt.Sprintf("Результат=%x", out))

	return result
}

func (c *StreamRC5) Decrypt(text string) []string {
	data, err := hex.DecodeString(text)
	if err != nil {
		return []string{"Нужен hex. Ошибка декодирования hex-строки"}
	}
	var result []string

	result = append(result, fmt.Sprintf("nonce=%08x", c.nonce))

	c.counter = 0
	var ks []byte

	out := make([]byte, len(data))
	for i := range data {
		if i%8 == 0 {
			ks = c.keystreamBlock()
			result = append(result, fmt.Sprintf("C[%d]=%x", c.counter-1, ks))
		}
		out[i] = data[i] ^ ks[i%8]
	}

	result = append(result, fmt.Sprintf("Результат=%s", out))

	return result
}

func (c *StreamRC5) Hack(text string) []string {
	var result []string
	result = append(result, "Тут ничего нет, а я думал должно быть, поэтому не стал переделывать архитектуру")

	return result
}

func (c *StreamRC5) keystreamBlock() []byte {
	block := make([]byte, 8)
	binary.LittleEndian.PutUint32(block[:4], c.nonce)
	binary.LittleEndian.PutUint32(block[4:], c.counter)
	c.counter++

	blockHex := fmt.Sprintf("%x", block)
	res := c.rc5.Encrypt(blockHex)
	ks, _ := hex.DecodeString(res[0][:16])
	return ks
}
