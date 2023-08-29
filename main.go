package main
 
import "fmt"

func main(){

	var conferenceName = "Go Programming"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// fmt.Println("Welcome to ",conferenceName," application")
	fmt.Printf("Welcome to %v application \n", conferenceName)
	// fmt.Println("We have total of ", conferenceTickets, " and available tickets of ",remainingTickets)
	fmt.Printf("We have total of %v  and available tickets of %v \n",conferenceTickets, remainingTickets)
	fmt.Println("Get your first programs starting here")
	fmt.Printf("conferenceTickets type is %T and conferenceName type is %T", conferenceTickets, conferenceName)
	fmt.Println(conferenceName)

	var bookings = [50] string{"barbie", "george", "nana"}
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var totalBookings = []string{}

	//Infinite loop
	for {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("No of tickets: ")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets-userTickets
		totalBookings = append(totalBookings, firstName+" "+lastName)

		fmt.Printf("Whole Array: %v \n", bookings)
		fmt.Println("first value of Array: \n", bookings[2])
		fmt.Printf("Array type : %T \n", bookings)
		fmt.Printf("Length of array:  %v \n", len(bookings))
		fmt.Printf("Thank you %v %v for booking %v tickets, you will recieve confirmation to email %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v Remaining tickets for %v \n", remainingTickets, conferenceName)
	}
}
	
	
	

	
	


	
