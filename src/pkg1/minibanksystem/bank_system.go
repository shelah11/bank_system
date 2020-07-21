package minibanksystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"crypto/rand"
	"log"
	"regexp"
	"strings"

	//"regexp"

)

type Testing struct {
	Name          string
	CurrentAmount float32
	Deposit       float32
	Withdraws     float32
	AccountNumber string
}

func (t *Testing) GetCurrentAmount() float32 {

	return t.CurrentAmount
}

func (t *Testing) AddIntoAccount(deposit float32) (float32, string) {
	t.CurrentAmount = t.CurrentAmount + deposit
	text := "deposited"
	return deposit, text
}

func (t *Testing) WithdrawsMoney(amount float32) (float32, error) {
	var err error

	if amount > t.CurrentAmount {
		err = errors.New("You do not have enough")
		return 0, err
	} else {
		t.Withdraws = amount
		t.CurrentAmount -= amount
	}

	return t.CurrentAmount, err
}

func IsLetter(input string) bool{
	customerName := strings.ToLower(input)

	var isLetter, _ = regexp.MatchString(`^[a-z]+$`,customerName )
	if !isLetter{
		return false
	}
	return true
}

func (t *Testing) OpenNewAccount(accountHolderName string, deposit float32) (string, error) {
	var err error

 	if len(accountHolderName) == 0{
		return "", errors.New("make sure you provide a valid name")
 	}
 	if IsLetter(accountHolderName) == false{
 		return "", errors.New("name must only contain letters")
	}
	if deposit == 0 || deposit < 20 {
		return "", errors.New("a minimum of £20 is required")
	}
	b := make([]byte, 4)
	_, err = rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	accountNumber := fmt.Sprintf("%d", b[0:4])

	newCustomer := Testing{
		Name:          accountHolderName,
		CurrentAmount: deposit,
		Deposit:       deposit,
		AccountNumber: accountNumber,
	}
	customer, err := json.MarshalIndent(newCustomer, "", " ")
	if err != nil{
		return "", err
	}
	return string(customer), err
}

//func accountReport(accountHolderName string, deposit float32) {
//	//var newUser map[string]float32
//	//newUser = make(map[string]float32)
//	//newUser[accountHolderName]deposit
//	//v, _ := uuid.NewV4()
//	//fmt.Printf("test uuid: %v", v)
//}
