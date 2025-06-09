package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets < remTickets

	if !isValidUserTickets || !isValidName || !isValidEmail {
		if !isValidName {
			fmt.Println("First Name or Last name you entered is too short. ")
		}
		if !isValidEmail {
			fmt.Println("You didnt enter @ in your email.")
		}
		if !isValidUserTickets {
			fmt.Println("Number of tickets you entered is invalid")
		}
		println("PLease enter valid credentials")
		return false
	} else {
		return true
	}
}
