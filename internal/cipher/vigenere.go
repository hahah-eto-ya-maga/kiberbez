package cipher

import (
	"fmt"
	"strings"
	"unicode"
)

type Vigenere struct {
	Key string
}

func NewVigenere(key string) *Vigenere {
	return &Vigenere{Key: key}
}

func (c *Vigenere) Name() string {
	return VigenereName
}

func (c *Vigenere) GetKey() string {
	return c.Key
}

func (c *Vigenere) SetKey(key any) {
	c.Key = key.(string)
}

func (c *Vigenere) Encrypt(text string) []string {
	var result []string
	result = append(result, encryptProcess(text, c.Key, false))
	return result
}

func (c *Vigenere) Decrypt(text string) []string {
	var result []string
	result = append(result, encryptProcess(text, c.Key, true))
	return result
}

func (c *Vigenere) Hack(text string) []string {
	var result []string

	textRunes := filterRussian(text)
	keyLen := findKeyLength(textRunes)

	originalDist := make([]float64, ABCSize)
	for r, f := range russianFreq {
		i := indexInABC(r)
		if i >= 0 {
			originalDist[i] = f
		}
	}

	var key []rune

	for i := 0; i < keyLen; i++ {
		var part []rune
		for j := i; j < len(textRunes); j += keyLen {
			part = append(part, textRunes[j])
		}

		dist := calculateDistribution(part)
		shift := calculateShift(originalDist, dist)

		key = append(key, ABC[shift])
	}

	plain := encryptProcess(text, string(key), true)

	result = append(result, fmt.Sprintf("Расшифрованный текст: %s\n", plain))
	result = append(result, fmt.Sprintf("Полученный ключ: %s", string(key)))

	return result
}

func buildKeyShifts(key string) []int {
	var result []int
	for _, r := range key {
		result = append(result, GetShift(r))
	}
	return result
}

func encryptProcess(text string, key string, invert bool) string {
	var result strings.Builder
	keyShifts := buildKeyShifts(key)
	keyIndex := 0

	for _, r := range text {
		idx := indexInABC(unicode.ToUpper(r))
		if idx == -1 {
			result.WriteRune(r)
			continue
		}

		shift := keyShifts[keyIndex]
		if invert {
			shift = -shift
		}

		result.WriteRune(ShiftRune(r, shift))

		keyIndex = (keyIndex + 1) % len(keyShifts)
	}

	return result.String()
}

var ABC = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
var ABCSize = len(ABC)
var russianFreq = map[rune]float64{
	'О': 0.1097, 'Е': 0.0845, 'А': 0.0801, 'И': 0.0735, 'Н': 0.0670,
	'Т': 0.0626, 'С': 0.0547, 'Р': 0.0473, 'В': 0.0454, 'Л': 0.0440,
	'К': 0.0349, 'М': 0.0321, 'Д': 0.0298, 'П': 0.0281, 'У': 0.0262,
	'Я': 0.0201, 'Ы': 0.0190, 'Ь': 0.0174, 'Г': 0.0170, 'З': 0.0165,
	'Б': 0.0159, 'Ч': 0.0144, 'Й': 0.0121, 'Х': 0.0097, 'Ж': 0.0094,
	'Ш': 0.0073, 'Ю': 0.0064, 'Ц': 0.0048, 'Щ': 0.0036, 'Э': 0.0032,
	'Ф': 0.0026, 'Ъ': 0.0004,
}

func indexInABC(r rune) int {
	for i, c := range ABC {
		if c == r {
			return i
		}
	}
	return -1
}

func calculateDistribution(text []rune) []float64 {
	dist := make([]float64, ABCSize)
	count := make([]int, ABCSize)

	for _, r := range text {
		idx := indexInABC(r)
		if idx >= 0 {
			count[idx]++
		}
	}

	for i := range dist {
		dist[i] = float64(count[i]) / float64(len(text))
	}

	return dist
}

func indexOfCoincidence(text []rune) float64 {
	dist := calculateDistribution(text)
	sum := 0.0
	for _, v := range dist {
		sum += v * v
	}
	return sum
}

func findKeyLength(text []rune) int {
	maxLen := 20
	ICs := make([]float64, maxLen+1)

	for k := 1; k <= maxLen; k++ {
		totalIC := 0.0
		for offset := 0; offset < k; offset++ {
			var group []rune
			for i := offset; i < len(text); i += k {
				group = append(group, text[i])
			}
			totalIC += indexOfCoincidence(group)
		}
		ICs[k] = totalIC / float64(k)
	}

	best := 1
	for k := 2; k < maxLen; k++ {
		if ICs[k] > ICs[k-1] && ICs[k] > ICs[k+1] {
			best = k
			break
		}
	}

	return best
}

func calculateShift(originalDist, dist []float64) int {
	bestShift := 0
	bestScore := -1.0

	for shift := 0; shift < ABCSize; shift++ {
		score := 0.0
		for i := 0; i < ABCSize; i++ {
			j := (i + shift) % ABCSize
			score += dist[j] * originalDist[i]
		}
		if score > bestScore {
			bestScore = score
			bestShift = shift
		}
	}

	return bestShift
}

func filterRussian(text string) []rune {
	var result []rune
	for _, r := range text {
		if indexInABC(unicode.ToUpper(r)) != -1 {
			result = append(result, unicode.ToUpper(r))
		}
	}
	return result
}
