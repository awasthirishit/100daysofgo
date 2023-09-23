package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	application := "app"

	//var bookings [50]string -> array
	var bookings []string //slice
	// bookings := []string {}

	fmt.Printf("Welcome to %v booking %v :)\n", conferenceName, application)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println("We have total of", remainingTickets, "tickets left!!!!")

	for {
		var firstname string
		var lastname string
		var email string
		var userTickets int

		fmt.Println("Enter your Firstname :")
		fmt.Scan(&firstname)

		fmt.Println("Enter your Lastname :")
		fmt.Scan(&lastname)

		fmt.Println("Enter your email address :")
		fmt.Scan(&email)

		fmt.Println("No. of tickets you want to book :")
		fmt.Scan(&userTickets)

		isValidName := len(firstname) >= 2 && len(lastname) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTicket := userTickets > 0 && userTickets > int(remainingTickets)

		if isValidName && isValidEmail && isValidUserTicket {
			remainingTickets = remainingTickets - uint(userTickets)
			//bookings[0] = firstname + " " + lastname
			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation mail on %v.\n", firstname, lastname, userTickets, email)
			fmt.Printf("%v tickets remaining for %v!!\n", remainingTickets, conferenceName)

			firstnames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}

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
