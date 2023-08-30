# Go Basics.

1. Installation
2. Getting Started
3. Variables
4. Formatting and fmt package
5. Arrays and Slices
6. Loops
7. if-else

* Download go for your machine on official website
https://go.dev/dl/

* Follow the installation steps and if you run go -version in terminal you should see something like below on successful installation: 
    ```go
    go version go1.21.0 darwin/amd64
    ```


* Install visual studio and Go extension in visual studio.


* Create a temp app called booking-app and open in visual studio
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
    and this creates a **go.mod** file which contains details like "module" and version of go that is used in our project.  This resolves the first error we encounter. 

* The first line of code is Print("hello world") 
* This shows a second error called "Declaration is required" which means go is telling us that it does not know where to start our application and an entry point of application is needed. 
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
Variables
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
* var v = 50 marks v as an integer value and var s = "abcd" marks s as a string value. 

* To define a type, 
    ```go
    var userName string
    ```

* %T is used to print the type of a variable
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

Arrays and Slices
* Arrays and Slices are commonly used datatypes in go applications. 

Arrays
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

Slices

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
  

Loops in Go

* Loops are used to execute code multiple times. 
* Loops in Go are simplified. We only have for loop. 
* Infinite loop is going to print hello world infinite number of times: 
  ```go 
  for{
    fmt.Println("Hello World")
  }
  
  ```

string.Fields()

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

Blank Identifier
* In the above example we will see an error for unused variable index. We did not use it in the for loop. We can use blank identifiers in go and this can be done with the help of underscore '_'
* Blank identifier is used to ignore a variable you dont want to use. 
* With go, you need to make unused variables explicit. 
  ```go
    firstNames := []string{}
	    for _, bookings := range bookings{
		    var names = string.Fields(bookings)
        firstNames = append(firstNames, names[0])
	    }
  ```

If-else
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
* 