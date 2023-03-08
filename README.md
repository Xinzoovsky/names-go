# Names-Go

Names-Go is a Go package that provides functionality for generating random first names, last names, and phone numbers.

## Installation

To install Names-Go, use the following command:

```sh
go get github.com/Xinzoovsky/names-go
```

## Usage

Here is an example of how to use Names Go to generate a random first name, last name, and phone number:

```go
package main

import (
	"fmt"
	"github.com/Xinzoovsky/names-go"
)

func main() {
	firstName := names.GetFirstName()
	lastName := names.GetLastName()
	phoneNumber := names.GetPhoneNumber()
    	email, _ := names.GetCatchall("@test.com", firstName, lastName)


	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Phone Number:", phoneNumber)
    	fmt.Println("Email:", email)
}
```
