package main

import (
	Account "demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := Account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}

	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *Account.Vault) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *Account.Vault) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func createAccount(vault *Account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := Account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логина")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(promt string) string {
	fmt.Println(promt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
