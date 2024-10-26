package Account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Account  []Account `json:"accounts"`
	UpdateAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Account:  []Account{},
			UpdateAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Account:  []Account{},
			UpdateAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Account {
		isMathed := strings.Contains(account.Url, url)
		if !isMathed {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Account = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Account {
		isMathed := strings.Contains(account.Url, url)
		if isMathed {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Account = append(vault.Account, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdateAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	files.WriteFile(data, "data.json")
}
