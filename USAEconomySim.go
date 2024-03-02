// USA Economic Game

package main

import "fmt"

var year int = 2047
var fed_funds_rate float32 = 25.00
var month int = 1
var month_name = month_names[0]
var month_names = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func main() {

	fmt.Print("Welcome to USA Economic Simulator \n\n")
	fmt.Print("You are playing as the chairman of the Federal Reserve \n")
	fmt.Print("The year is 2047 and rates are at an all time high of 25%\n")
	fmt.Print("The consumer is struggling and the country is looking for some hope out of this economic downturn\n")

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

	fmt.Print("\n\nIt is now ", month_name, " ", year)

	fmt.Print("\n\nOptions: \n")
	fmt.Print("1. Change Federal Funds Rate\n")
	fmt.Print("2. Advance a month\n")
	fmt.Print("\nEnter your choice: ")

	fmt.Scan(&choice)

	switch choice {

	case 1:
		changeFedFundsRate()

	case 2:
		monthChange()

	}

}

func changeFedFundsRate() {

	fmt.Print("Current Federal Funds Rate: ", fed_funds_rate, "%")
	fmt.Print("\n\nOptions:\n")
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

func monthChange() {

	month = month + 1

	switch month {
	case 1:
		month_name = month_names[0]

	case 2:
		month_name = month_names[1]

	case 3:
		month_name = month_names[2]

	case 4:
		month_name = month_names[3]

	case 5:
		month_name = month_names[4]

	case 6:
		month_name = month_names[5]

	case 7:
		month_name = month_names[6]

	case 8:
		month_name = month_names[7]

	case 9:
		month_name = month_names[8]

	case 10:
		month_name = month_names[9]

	case 11:
		month_name = month_names[10]

	case 12:
		month_name = month_names[11]
	}

	if month == 13 { // bug to fix the month and year once time goes from December to January
		month = month - 12
		yearChange()
	}

	options()

}

func lowerFedFundsRate() {

	fmt.Print("Enter how much you want to lower the Federal Funds Rate: ")
	var lower_amount float32
	fmt.Scan(&lower_amount)

	fed_funds_rate = fed_funds_rate - lower_amount
	fmt.Print("The new Federal Funds Rate is ", fed_funds_rate, "%")

	options()

}

func raiseFedFundsRate() {

	fmt.Print("Enter how much you want to raise the Federal Funds Rate: ")
	var raise_amount float32
	fmt.Scan(&raise_amount)

	fed_funds_rate = fed_funds_rate + raise_amount
	fmt.Print("The new Federal Funds Rate is ", fed_funds_rate, "%")

	options()

}
