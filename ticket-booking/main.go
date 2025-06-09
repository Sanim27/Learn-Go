package main

import (
	"fmt"
	"learning_go/helper"
	"sync"
	"time"
)

var confName string = "Go Conference" //only applies to variables and cant explicitly tell dataytype
const confTickets uint = 50

var remTickets uint = 50
var bookings = make([]UserData, 0) //if there is no number inside [] then it is a slice.

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(confName, remTickets, confTickets)

	firstName, lastName, email, userTickets := getUserInput()

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	if userTickets > remTickets {
		fmt.Printf("We only have %v tickets remaining. You cant book %v tickets\n", remTickets, userTickets)
		//continue
	}

	isUserInputValid := helper.ValidateUserInput(firstName, lastName, email, userTickets, remTickets)
	if !isUserInputValid {
		//break
	}

	remTickets = remTickets - userTickets
	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is : %v \n", bookings)

	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v number of tickets remaining for %v\n", remTickets, confName)

	wg.Add(1)
	go sendTicket(firstName, userTickets, email)

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("First name of users who booked:  %v\n", firstNames)

	if remTickets == 0 {
		//end program
		fmt.Println("Our Conference is booked out. Come back next year.")
		//break
	}
	wg.Wait()

}

func greetUsers(confName string, remTickets uint, confTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for inputs

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want to book")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func sendTicket(firstName string, userTickets uint, email string) {
	time.Sleep(time.Second * 10)
	var Ticket = fmt.Sprintf("%v number of ticket for %v\n", userTickets, firstName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", Ticket, email)
	fmt.Println("###################")
	wg.Done()
}
