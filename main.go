package main
 
import ("fmt" 
		"strings"
)

func main(){

	
	var conferenceName = "Go Programming"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	// fmt.Println("Welcome to ",conferenceName," application")
	// fmt.Println("We have total of ", conferenceTickets, " and available tickets of ",remainingTickets)
	// fmt.Printf("\n conferenceTickets type is %T and conferenceName type is %T", conferenceTickets, conferenceName)
	// var bookings = [50] string{"barbie", "george", "nana"}
	
	var totalBookings = []string{}
	// Conditional for loop
	// for remainingTickets>0 && len(totalBookings)<50{}
	//Infinite loop
	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketCount:= validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if(isValidName && isValidEmail && isValidTicketCount){
			bookTicket(remainingTickets, userTickets , totalBookings , firstName , lastName , email , conferenceName){

			// call function print firstnames
			firstNames := printFirstNames(totalBookings)
			fmt.Printf("\nFirst names %v:", firstNames)
	
			// fmt.Printf("The first names of bookings are %v \n", firstNames)
	
			if remainingTickets == 0{
				fmt.Println("Tickets are sold out!!")
				break
			}
			
		}else{
			if !isValidName{
				fmt.Println("too short name")
			}
			if !isValidEmail{
				fmt.Println("Email is missing @")
			}
			if(!isValidTicketCount){
				fmt.Println("Invalid ticket count")
			}
		}
	}
}

func greetUsers(conferenceName string, confTickets int,remainingTickets uint){
	fmt.Printf("Welcome to our Learning path %v", conferenceName)
	fmt.Printf("\nWe have total of %v  and available tickets of %v \n",confTickets, remainingTickets)
	fmt.Println("\nGet your first programs starting here")
	fmt.Println(conferenceName)
}

func printFirstNames(totalBookings []string) []string{
	firstNames := []string{}
	for _, totalBookings := range totalBookings{
		var names = strings.Fields(totalBookings)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName)>=2 && len(lastName)>=2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketCount = userTickets>0 && remainingTickets>=userTickets

	return isValidName, isValidEmail, isValidTicketCount

}

func getUserInput() (string, string, string, uint){
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

func bookTicket(remainingTickets uint, userTickets uint, totalBookings []string, firstName string, lastName string, email string, conferenceName string){
	remainingTickets = remainingTickets-userTickets
	totalBookings = append(totalBookings, firstName+" "+lastName)	
			// fmt.Printf("Whole Array: %v \n", totalBookings)
			// fmt.Println("first value of Array: \n", bookings[2])
			// fmt.Printf("Array type : %T \n", bookings)
			// fmt.Printf("Length of array:  %v \n", len(bookings))
	fmt.Printf("\nThank you %v %v for booking %v tickets, you will recieve confirmation to email %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v Remaining tickets for %v \n", remainingTickets, conferenceName)
	
}