# Go Basics.

1. Installation
2. Getting Started
3. Variables
4. Formatting and fmt package
5. Arrays and Slices
6. Loops
7. Blank Identifier
8. if-else
9. Input validations
10. Switch statement
11. Functions
12. Package level variables
13. Local Variables
14. Packages 
15. Maps
16. Struct

**Installation**

* Download go for your machine on official website
https://go.dev/dl/

* Follow the installation steps and if you run `go -version` in terminal you should see something like below on successful installation: 
    ```go
    go version go1.21.0 darwin/amd64
    ```


* Install visual studio and Go extension in visual studio.


* Create a temp app called `booking-app` and open in visual studio
    ```console
    cmd: mkdir booking-app 
    ```
* Create a new file called main.go
    * First error we see is go.mod file not found 
    * In order to fix it
    * Way 1: go to terminal and create mod file. In order to do that run the command 
    ```go
    go mod init booking-app
    ```
    * Way 2: Go to terminal and disable modules
     ```go
    go env -w GO111MODULE=off
    ```
    

* In our example we are going to add a mod file using  
    ```go
    go mod init booking-app
    ``` 
    and this creates a `go.mod` file which contains details like "module" and version of go that is used in our project.  This resolves the first error we encounter. 

* The first line of code is `Print("hello world")` 
* This shows a second error called `Declaration is required` which means go is telling us that it does not know where to start our application and an entry point of application is needed. 
We always need a main starting point for execution. In order to do that, we need a main() function that go will look for while executing our main application. 

    ```go 
    func main(){}
    ``````
* Once this error is resolved, we see a new error called print is not found. Print comes from a go package called fmt and we have to specify any imports that use those functionality. The followig command resolves the error

    ```go 
    import "fmt"
    ```

* Think of packages as containers of various functionalities that go gives us to readily use them. You can find documentation about packages here:
https://pkg.go.dev/std

* Now as all the errors are resolved, we can run our first go application by 
    ```go
    go run main.go
    `````` 
    which outputs 
     ```
    Hello World
    ``````

**Variables**
* Variables are used to store values
* Give the variable a name and reference everywhere in the app
* Update value only once
* Make our app more flexible
* Variables in go can be created using var keyword
    ```go
    var conferenceName = "Go Programming"
    ```
* for variables that do not change, we use const
more like final keyword in java. 
    ```go
    const conferenceTickets = 50
    ```
* If we try to update constant values by modifying them, we get an error message saying "Cannot declare to [Variable]"

* We can do formatting of various strings using printf package of go. 
    ```go
    fmt.Printf("Welcome to %v application \n", conferenceName)


    fmt.Printf("We have total of %v  and available tickets of %v \n",conferenceTickets, remainingTickets)
    ```

* Go is a statically typed language. So if a variable is not initialized, we should mention its type. Like if it is a string/integer value. If the value is initialized, go interpreter knows the type of variable that is being passed. 
* `var v = 50` marks v as an integer value and `var s = "abcd"` marks s as a string value. 

* To define a type, 
    ```go
    var userName string
    ```

* `%T `is used to print the type of a variable
  ```go
  fmt.Printf("conferenceTickets type is %T and conferenceName type is %T", conferenceTickets, conferenceName)
  ```

  We can also create variables by using the command :=
  ```go
  test := "Hello World"
  ```

* fmt.Scan is used for scanning the user input
    ```go
    fmt.Scan(&userName)
    ```
* & is used as a pointer. When we create a variable, the value gets stored on memory fo our computer. Whenever we reference a variable name go compiler checks the memory where the variable is stored/Memory address. Pointer is the variable that points to memory address of another variable which points to actual value. 
* They are also called special variables. 
  ```go
  fmt.Println(&userName)
  ``` 
  gives the hash/pointer to the memory locations. 
* We cannot use variables before the actual declaration like javascript. Order matters in go 

**Arrays and Slices**
* Arrays and Slices are commonly used datatypes in go applications. 

**Arrays**
* Arrays in go have fixed size. 
  ```go
  var variableName [size]variable_type
    ```
* We cannot mix different datatypes in one array. Something like : 
  ```go 
  var bookings = [50] string{"barbie", "george", 3, 5}
  ```
* We can add elements to array based on index position. 
  ```go 
  bookings[0]="Barbie"
  bookings[9]="Crime Master Go Go"
  ```

**Slices**

* Slices is an abstraction of array 
* Slices are more flexible and powerful. They have variable length or get a subarray of its own. 
* Slices are also index-based and have a size but can be resized when needed. 
* For defining a slice, we create an array without size definition. 
  ```go
  var bookings []string OR
  var bookings = []string{} OR
  bookings := []string{}
  ```
* We can add elements to slice  using "append()" a builtin function in go
  ```go 
  bookings = append(bookings, "Crime Master GO GO")
  ```
* Advantage of slices is there is no need to keep track of indexes and we can allocate information dynamically. Which can save memory and space. 
  

**Loops in Go**

* Loops are used to execute code multiple times. 
* Loops in Go are simplified. We only have for loop. 
* Infinite loop is going to print hello world infinite number of times: 
  ```go 
  for{
    fmt.Println("Hello World")
  }
  
  ```

**string.Fields()**

* string.Fields() splits the string with white space separator and returns the slice with split elements. 
```go
  firstNames := []string{}
	  for index, bookings := range bookings{
		  var names = string.Fields(bookings)
      firstNames = append(firstNames, names[0])
	  }
```
* To use string functions like Fields we need to import the package "strings". It is important to declare imported packages in new lines and brackets should be declared as shown below: 
  ```go 
  import ("fmt" 
		"strings"
  )
  ```

**Blank Identifier**
* In the above example we will see an error for unused variable index. We did not use it in the for loop. We can use blank identifiers in go and this can be done with the help of underscore '_'
* Blank identifier is used to ignore a variable you dont want to use. 
* With go, you need to make unused variables explicit. 
  ```go
  firstNames := []string{}
	    for _, bookings := range bookings{
		    var names = string.Fields(bookings)
        // Here names will have 2 values firstName and lastName. 
        // names[0] returns the firstname
        firstNames = append(firstNames, names[0])
	    }
  ```

**If-else**
* This works similar to how if-else works in java. 
* It is important to input correct conditional statement in go. If a condition fails, that should be handled by user or else you might endup seeing redundant data. For example, we have 50 tickets and we try to book 54 tickets with the below condition: 
  ```go
  if remainingTickets == 0{
			fmt.Println("Tickets are sold out!!")
			break
		}

    OUTPUT: 
    18446744073709551610 Remaining tickets for Go Programming 
  ```
* In order to avoid such random outputs and infinite loops, we have to ensure that right conditional statements are in place. 
* We can use break and continue similar to how we use in java.

Conditional for loop
* Execute a loop till a condition is satisfied
  ```go
  for remainingTickets>0 && len(totalBookings)<50{ 
    // do something
  }
  ```

**Input validations**
* We need to ensure that application can handle bad user input data
* We can do string validations using "string" package. 
* len() calculates the length of a string.
* strings.Contains(val1, val2) can be used to make string validations 
  ```go
    //Here isvalidName is a boolean value that returns true if length of both first and last name is >2
    var isValidName = len(firstName)>=2 && len(lastName)>=2
    //Here isValidEmail is a boolean value that returns true if email contains the character @
    var isValidEmail = strings.Contains(email, "@")
    //Here isValidTicketCount is a boolean value that returns true if conditional statement passes for the integer comparisions. 
    var isValidTicketCount = userTickets>0 && remainingTickets>=userTickets
  ```

**Switch statement**
* works similar to java switch case. Example below: 
  ```go
  city := "London"

  switch{
    case "New York":
	    //Execute code for reservations in New York
    case "London":
	    //Execute code for reservations in London
    case "Singapore":
	    //Execute code for reservations in Singapore
    case "Mumbai":
	    //Execute code for reservations in Mumbai
    default:
	    fmt.Println("Invalid city selected")
  å}```

**Functions**
* We encapsulate code into its own container called function. Which logically belong together. 
* Like variable name you have to give a function a descriptive name. 
* Call the function by its name, whenever you want to execute this code. 
  ```go
  func greetUsers(){
	  fmt.Println("Welcome to our Learning path")
   }
  ```
* Creating/declaring a function does not gaurantee that it will be executed. This only means that the created function can be for later use.
* Functions will only be executed when called. 
* Functions can be called as many times as you want. 
* Functions allow us to reduce code duplication. 
* We can call a function by mentioning it in main function 
  ```go
  func main(){
    //your code
    greetUsers()
  }
  ```
* We can also have functions that accept parameters. We can do that by mentioning parameter and type in the function declaration
  ```go
  func greetUsers(confName string){
	  fmt.Printf("Welcome to our %v", confName)
   }
  ```
* Returning values from a function 
* A function can return data as a result
* Function can take an input and return an output. 
* In Go you have to define the input and output parameters including its type explicitly. 
* We have the return type of the function outside of the function delcaration brackets. 
  ```go
  func printFirstNames(totalBookings []string) []string{
	    firstNames := []string{}
	    for _, totalBookings := range totalBookings{
		    var names = strings.Fields(totalBookings)
		    var firstName = names[0]
		    firstNames = append(firstNames, firstName)
	    }
	  return firstNames
    }
    ```
* Go has the capability to return multiple values for a function. 
* We need to define multiple return types in paranthesis in the function declaration
  i.e., input parameters in (val1, val2, ..) and output types in second paranthesis as shown below: 
  ```go
  func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	  var isValidName = len(firstName)>=2 && len(lastName)>=2
	  var isValidEmail = strings.Contains(email, "@")
	  var isValidTicketCount = userTickets>0 && remainingTickets>=userTickets
	  return isValidName, isValidEmail, isValidTicketCount
  }
  ```
* While calling this function with multiple return types in main() function, we can declare multiple variables and assign it to single function and the return values are saved in that order. 
  ```go
  isValidName, isValidEmail, isValidTicketCount:= validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

  ```
More use cases of functions: 
* Group logic that belongs together. 
* Reuse logic and reduce duplication of code

**Package level variables**
* Package level variables are declared at the top outside all functions
* They can be accessed inside any function in the package and in all files which are in same package. 
* Package level variables must be declared using var keyword. They cannot be declared with :=
  ```go
  package main
  import (
	  "fmt"
	  "strings"
    )
    // Package level variables
    const conferenceTickets = 50
    var conferenceName = "Go Programming"
    var remainingTickets uint = 50
    var totalBookings = []string{}
    func main() {//Your code
    }

  ```

**Local Variables**
* Local variables are declared inside a function or block
* They can be accessed only inside a function or a block.

Packages 
* Can I modularize an application in go? 
* Absolutely, go programs are organized into packages. 
* A package is a collection of go files which can be (1..n) files. 
* At the Scope of package level, variables and fucntions defined outside any function can be accessed by all other files within same package. 
* In our example we added a new helper package called helper.go
* When you have methods/function from one package being refered in another., when we run  ```go run main.go``` we will see an error message ```./main.go:28:52: undefined: validateUserInput``` Inorder to overcome errors like this, we will need to call all the files that belong to the application. 
  ```go
  go run main.go helper.go OR 
  go run .
  ```
* "." specifies the current folder and all the files in that folder will be picked during execution. 
* Multiple Packages can help you organize your application and logically group your code. 
* When you have to call function from a different package, similar to how we use "strings" or "fmt" we have to explicitly import our user defined packages to use the methods. 
* In our example we created a folder called helper, with a helper package. Can we import it directly by `import "helper"` ? No we cannot. 
* In order to import userdefined packages, we need to specify the path. We have to checkthe go.mod file for the import path. In this example `module booking-app` defines the import path. 
  ```go 
  import (
	    "fmt"
	    "strings"
	    "booking-app/helper"
    )
  ```
* For a function to be used between packages, we need to explicitly export that function. 
* In go to export a function, capitalize the first alphabet of a function and it is available for exporting. 
  ```go
  func ValidateUserInput(//Your code
    )
    and you can use it in other packages using `helper.ValidateUserInput``
  ```
* One more thing you need to notice here is all the `fmt` or `strings` package functions start with capital letters. 
* Along with functions, we can also export variables in go. 
* We have 3 levels of scope
  * Local 
  1. Declared within function and can be used only within that function.
  2. Declared with in block like if-else or for and can be used only in that block. 
  * Package level
   Declared outside all functions and can be used everywhere in same package. 
  * Global
    Variables that are shared across packages are called global variables. Declared outside all functions and uppercase first letter. They can be used everywhere across all packages. 
* Variable scope is defined as the region of program where a defined variable can be accessed. 

**Maps**
* A datatype that allows us to store multiple key-value pairs is called a map. 
* Maps have unique keys to values. 
  ```go 
    var userData = make(map[string]string)
  ```
* We cannot mix datatypes in go. 
* To convert an integer to string we have a function called `FormatUint` which comes from a package called `strconv` and we need to import "strconv" to use the package. 
  ```go
  strconv.FormatUint(uint64(userTickets), 10)
  ```
**Struct Datatype**
* Struct is used when we need to use values of mixed datatypes. 
* Maps in go limit to support only 1 datatype and we have to use type conversions as shown in example above to support different datatypes. 
* Struct stands for structures. 
* Struct allows us to save key value pairs but with mixed datatypes. 
* Struct should be defined with `type` keyword with `<name of struct>` and `struct` keyword. Inside the curly braces we need to define the variables along with the types. 
  ```go
  type UserData struct {
	fName       string
	lName       string
	mail        string
	noOfTickets uint
  }
  ```
* `type` keyword creates a new type with the name you specify. 
* You could also create a type based on every other data type like int, string, uint etc., 
* Struct is similar to classes in java which do not support inheretence. 
* Usage in function: 
 ```go 
  var userData = UserData{
	  fName:       firstName,
		lName:       lastName,
		mail:        email,
		noOfTickets: userTickets,
	}
  var totalBookings = make([]UserData, 0)
  totalBookings = append(totalBookings, userData)
 ```

Concurrency 
