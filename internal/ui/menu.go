package ui

import (
	"bufio"
	"fmt"
	"kiberbez/internal/ui/colors"
	"kiberbez/internal/ui/menuActions"
	"kiberbez/internal/ui/utils"
	"os"
)

type MenuAction struct {
	Title string
	Exec  func()
}

func RunMenu() {
	reader := bufio.NewReader(os.Stdin)

	utils.ClearScreen()
	c := SelectCipher()
	utils.ClearScreen()

	for {
		fmt.Println("[Шифр] " + colors.GREEN + c.Name() + colors.DEFAULT)
		fmt.Println("[Ключ] " + colors.GREEN + c.GetKey() + colors.DEFAULT)

		actions := []MenuAction{
			{"Зашифровать", func() {
				menuActions.EncryptCase(reader, c)
			}},
			{"Дешифровать", func() {
				menuActions.DecryptCase(reader, c)
			}},
			{"Взломать", func() {
				menuActions.HackCase(reader, c)
			}},
			{"Поменять ключ", func() {
				menuActions.ReadKeyForCipherCase(reader, c)
			}},
			{"Поменять шифр", func() {
				c = SelectCipher()
			}},
			{"Очистить экран", func() {
				utils.ClearScreen()
			}},
			{"Выйти", func() {
				os.Exit(0)
			}},
		}

		choice := utils.GetUserChoice(actions, func(action MenuAction) {
			fmt.Print(action.Title)
		})

		choice.Exec()
		fmt.Println()
		if choice.Title != "Очистить экран" {
			fmt.Println("--------------------------------------------------")
		}
		fmt.Println()
	}
}
