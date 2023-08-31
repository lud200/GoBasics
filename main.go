package main

import (
	// "booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Package level variables
const conferenceTickets = 50

var conferenceName = "Go Programming"
var remainingTickets uint = 50

// Initializing a list of maps:
// var totalBookings = make([]map[string]string, 0)

type UserData struct {
	fName       string
	lName       string
	mail        string
	noOfTickets uint
}

var totalBookings = make([]UserData, 0)

var wait = sync.WaitGroup{}

func main() {

	greetUsers()
	// fmt.Println("Welcome to ",conferenceName," application")
	// fmt.Println("We have total of ", conferenceTickets, " and available tickets of ",remainingTickets)
	// fmt.Printf("\n conferenceTickets type is %T and conferenceName type is %T", conferenceTickets, conferenceName)
	// var bookings = [50] string{"barbie", "george", "nana"}
	// Conditional for loop
	// for remainingTickets>0 && len(totalBookings)<50{}
	//Infinite loop
	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketCount {
		bookTicket(userTickets, firstName, lastName, email)

		wait.Add(1)

		go sendTicket(userTickets, firstName, lastName, email)

		// call function print firstnames
		firstNames := printFirstNames()
		fmt.Printf("\nFirst names %v:", firstNames)

		// fmt.Printf("The first names of bookings are %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Tickets are sold out!!")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("too short name")
		}
		if !isValidEmail {
			fmt.Println("Email is missing @")
		}
		if !isValidTicketCount {
			fmt.Println("Invalid ticket count")
		}
	}
	wait.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to our Learning path %v", conferenceName)
	fmt.Printf("\nWe have total of %v  and available tickets of %v \n", conferenceTickets, remainingTickets)
	fmt.Println("\nGet your first programs starting here")
	fmt.Println(conferenceName)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, totalBookings := range totalBookings {
		// var names = strings.Fields(totalBookings)
		// var firstName = names[0]
		// firstNames = append(firstNames, totalBookings["firstName"])
		firstNames = append(firstNames, totalBookings.fName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("No of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for user
	// var userData = make(map[string]string)
	var userData = UserData{
		fName:       firstName,
		lName:       lastName,
		mail:        email,
		noOfTickets: userTickets,
	}
	totalBookings = append(totalBookings, userData)
	fmt.Printf("Total bookings = %v", totalBookings)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// fmt.Printf("Whole Array: %v \n", totalBookings)
	// fmt.Println("first value of Array: \n", bookings[2])
	// fmt.Printf("Array type : %T \n", bookings)
	// fmt.Printf("Length of array:  %v \n", len(bookings))
	fmt.Printf("\nThank you %v %v for booking %v tickets, you will recieve confirmation to email %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v Remaining tickets for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("***********")
	fmt.Printf("Sending ticket: %v \n to email %v\n", ticket, email)
	fmt.Println("***********")
	wait.Done()
}
