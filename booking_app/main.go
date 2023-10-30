package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string //slice

func main() {

	//var bookings [50]string -> array

	// bookings := []string {}

	greetUsers()

	for {

		firstname, lastname, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTicket := validateUserInput(firstname, lastname, email, userTickets)

		if isValidName && isValidEmail && isValidUserTicket {
			bookTickets(firstname, lastname, email, userTickets)

			firstnames := getFirstNames()
			fmt.Printf("Total Bookings %v\n", firstnames)

			if remainingTickets == 0 {
				fmt.Printf("Sold out!! Comeback next year :(")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("You name is less than two digits")
			} else if !isValidEmail {
				fmt.Println("Your Email is missing @")
			} else if !isValidUserTicket {
				fmt.Println("Number of tickets is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application :)\n", conferenceName)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println("We have total of", remainingTickets, "tickets left!!!!")
}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	fmt.Println("Enter your Firstname :")
	fmt.Scan(&firstname)

	fmt.Println("Enter your Lastname :")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address :")
	fmt.Scan(&email)

	fmt.Println("No. of tickets you want to book :")
	fmt.Scan(&userTickets)

	return firstname, lastname, email, userTickets
}

func validateUserInput(firstname string, lastname string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets < remainingTickets

	return isValidName, isValidEmail, isValidUserTicket
}

func bookTickets(firstname string, lastname string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstname + " " + lastname
	bookings = append(bookings, firstname+" "+lastname)

	fmt.Printf("Hi! , Thank you %v %v for booking %v tickets. You will recieve a confirmation mail on %v.\n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v!!\n", remainingTickets, conferenceName)

}

func getFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames
}
