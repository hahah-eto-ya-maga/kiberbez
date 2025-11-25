package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunMenu() {
	reader := bufio.NewReader(os.Stdin)

	c := SelectCipher()

	for {
		fmt.Println("\n---", c.Name(), "Cipher ---")
		fmt.Println("1. Encrypt")
		fmt.Println("2. Decrypt")
		fmt.Println("3. Bruteforce")
		fmt.Println("9. Change cipher")
		fmt.Println("0. Exit")

		fmt.Print("Choose: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			fmt.Println("Result:", c.Encrypt(text))

		case "2":
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			fmt.Println("Result:", c.Decrypt(text))

		case "3":
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			fmt.Println("Results:")
			for _, r := range c.Bruteforce(text) {
				fmt.Println(" -", r)
			}

		case "9":
			c = SelectCipher()

		case "0":
			return
		}
	}
}
