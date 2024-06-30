package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, ticketCounts uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValideEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticketCounts > 0 && ticketCounts <= remainingTickets
	return isValidName, isValideEmail, isValidTicketNumber
}