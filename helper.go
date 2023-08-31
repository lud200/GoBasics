package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketCount = userTickets > 0 && remainingTickets >= userTickets

	return isValidName, isValidEmail, isValidTicketCount

}
