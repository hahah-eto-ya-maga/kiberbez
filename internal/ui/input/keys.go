package input

import (
	"bufio"
	"fmt"
	"kiberbez/internal/cipher"
	"math/big"
	"strconv"
	"strings"
)

func ReadKeyForCipher(reader *bufio.Reader, cipherName string) any {
	switch strings.ToLower(cipherName) {
	case cipher.CaesarName:
		for {
			fmt.Print("Введите ключ количество граней (целое число): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			k, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Неверный ввод, повторите попытку")
				continue
			}
			return k
		}
	case cipher.VigenereName:
		for {
			fmt.Print("Введите ключ слово (строка из БУКВ): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if !IsAlphaCyrillic(input) {
				fmt.Println("Неверный ввод, повторите попытку. Только кириллица")
				continue
			}

			return input
		}
	case cipher.RC5Name, cipher.StreamRC5Name:
		for {
			fmt.Print("Введите ключ: ")
			input, _ := reader.ReadString('\n')
			key := strings.TrimSpace(input)

			var rounds int
			for {
				fmt.Print("Введите количество раундов (от 0 до 255): ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				var err error
				rounds, err = strconv.Atoi(input)
				if err != nil || rounds < 0 || rounds > 255 {
					fmt.Println("Неверный ввод, повторите попытку")
					continue
				}
				break
			}

			return cipher.RC5Key{
				Key:    []byte(key),
				Rounds: rounds,
			}
		}
	case cipher.RSAName:
		for {
			fmt.Print("Введите p (простое число): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			px, err := strconv.Atoi(input)
			if err != nil || !big.NewInt(int64(px)).ProbablyPrime(20) {
				fmt.Println("Неверный ввод, повторите попытку. p должно быть простым")
				continue
			}

			fmt.Print("Введите q (простое число): ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			qx, err := strconv.Atoi(input)
			if err != nil || !big.NewInt(int64(qx)).ProbablyPrime(20) {
				fmt.Println("Неверный ввод, повторите попытку. q должно быть простым")
				continue
			}

			fmt.Print("Введите e: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			ex, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Неверный ввод, повторите попытку")
				continue
			}

			p := big.NewInt(int64(px))
			q := big.NewInt(int64(qx))
			e := big.NewInt(int64(ex))

			phi := new(big.Int).Mul(
				new(big.Int).Sub(p, big.NewInt(1)),
				new(big.Int).Sub(q, big.NewInt(1)),
			)

			gcd := new(big.Int)
			gcd.GCD(nil, nil, e, phi)
			if gcd.Cmp(big.NewInt(1)) != 0 {
				fmt.Println("e должно быть взаимно простым с фи(n) (функция Эйлера), повторите попытку")
				continue
			}

			return cipher.RSAProps{
				P: p,
				Q: q,
				E: e,
			}
		}
	}

	return nil
}
