package main

import (
	"banksystem/src/pkg1/minibanksystem"
	"fmt"
)
func main() {
	var name string
	var deposit float32
	var withdraw float32
	var option int

	first := minibanksystem.Testing{}

	fmt.Println("Press 1 to open an account")
	fmt.Println("Press 2 to check balance")
	fmt.Println("Press 3 to deposit into account")
	fmt.Println("Press 4 withdraw from account")
	fmt.Scanln(&option)

	switch {
	case option == 1:
		fmt.Print("Enter your full name: ")
		fmt.Scanln(&name)

		fmt.Print("How much would you like to deposit: ")
		fmt.Scanln(&deposit)
		firstCustomer, _ := first.OpenNewAccount(name, deposit)
		fmt.Printf("%v", firstCustomer)
	case option == 2:
		first.GetCurrentAmount()
	case option == 3:
		fmt.Print("How much would you like to deposit: ")
		fmt.Scanln(&deposit)
		first.AddIntoAccount(deposit)
	case option == 4:
		fmt.Print("How much would you like to deposit: ")
		fmt.Scanln(&withdraw)
		first.WithdrawsMoney(withdraw)
	}

}

