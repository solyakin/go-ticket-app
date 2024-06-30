package main

import (
	"fmt"
	"sync"
	"time"
)

var ticketName string = "Summer Jam"
var remainingTickets uint = 50
var bookings = make([]UserType, 0)

type UserType struct {
	firstName string
	lastName string
	email string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main () {

	greetUser(ticketName, remainingTickets)

	// for {

		firstName, lastName, email, ticketCounts := getUserInput()
		isValidName, isValideEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, ticketCounts)

		if isValidName && isValideEmail && isValidTicketNumber {

			bookTicket(ticketCounts, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(ticketCounts, firstName, lastName, email)
			firstNames := getFirstName()
			fmt.Printf("These are all the bookings: %v\n", firstNames)
			
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out, come back next year.")
				// break
			}
		}else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValideEmail {
				fmt.Println("email address is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("ticket number is invalid")
			}
		}

		wg.Wait()
	// }

	//switch statement

	city := "London"
	switch city {
		case "New York": 
			// eexcute code for booking new york ticket
		case "London": 
			// eexcute code for booking new york ticket
		case "Berlin": 
			// eexcute code for booking new york ticket
		default:
			fmt.Printf("No valid city selected")
	}
}

func greetUser(conferenceName string, ticketCount uint) {
	fmt.Printf("WELCOME TO %v BOOKING TICKETS\n", conferenceName)
	fmt.Printf("We have %v ticket availabe\n", ticketCount)
}

func bookTicket( ticketCounts uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - ticketCounts

	//create a map for a user
	var userData = UserType {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberofTickets: ticketCounts,
	}

	bookings = append(bookings, userData)

	//slice
	// bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, ticketCounts, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ticketName)
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var ticketCounts uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&ticketCounts)

	return firstName, lastName, email, ticketCounts
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
