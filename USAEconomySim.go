// USA Economic Game

package main

import "fmt"

var year int = 2047
var fed_funds_rate float32 = 25.00

func main() {

	fmt.Print("Welcome to USA Economic Simulator \n\n")
	fmt.Print("You are playing as the chairman of the Federal Reserve \n")
	fmt.Print("The year is 2047\n")

	var player_first_name, player_last_name string

	fmt.Print("Enter a first name for yourself, chairman: ")
	fmt.Scan(&player_first_name)
	fmt.Print("Enter a last name for yourself, chairman: ")
	fmt.Scan(&player_last_name)

	fmt.Print("Welcome chairman " + player_first_name + " " + player_last_name)

	options()
}

func options() {

	var choice int

	fmt.Print("\n\nOptions: \n")
	fmt.Print("1. Change Federal Funds Rate\n")
	fmt.Print("\nEnter your choice: ")

	fmt.Scan(&choice)

	switch choice {

	case 1:
		changeFedFundsRate()

	}

}

func changeFedFundsRate() {

	fmt.Print("Current Federal Funds Rate: ", fed_funds_rate)
	fmt.Print("\nOptions:\n")
	fmt.Print("1. Raise the Federal Funds Rate")
	fmt.Print("\n2. Lower the Federal Funds Rate")
	fmt.Print("\n3. Keep the Federal Funds Rate the same")

	var fed_funds_question int

	fmt.Print("\nEnter your choice: ")

	fmt.Scan(&fed_funds_question)

	switch fed_funds_question {

	case 1:
		raiseFedFundsRate()

	case 2:
		lowerFedFundsRate()

	case 3:
		options()

	}

}

func yearChange() {

	year = year + 1
	fmt.Print("The year is now ", year)

}

func lowerFedFundsRate() {

	fmt.Print("Enter how much you want to lower the Federal Funds Rate: ")
	var lower_amount float32
	fmt.Scan(&lower_amount)

	fed_funds_rate = fed_funds_rate - lower_amount
	fmt.Print("The new Federal Funds Rate is ", fed_funds_rate)

}

func raiseFedFundsRate() {

	fmt.Print("Enter how much you want to raise the Federal Funds Rate: ")
	var raise_amount float32
	fmt.Scan(&raise_amount)

	fed_funds_rate = fed_funds_rate + raise_amount
	fmt.Print("The new Federal Funds Rate is ", fed_funds_rate)

}
