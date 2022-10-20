# Kopeechka Wrapper (Golang)

# Functions:
* Get Balance
* Buy Email 
* Cancel Email 
* Get Letter 

### Full Docuemtation Of Client Below!

**Usage Below**
```go 
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ItsMerkz/Kopeechka_Wrapper/"
)

func main() {
	// Email Client!
	var emailkey string = "" // Client Key

	client := new(kopeechka.EmailClient)

	// Buys Email Taking 3 arguements, returning mailid and email
	MailId, Email := client.BuyEmail(emailkey, "Domain", "Target_Url")
	fmt.Printf("%v | %v\n", MailId, Email) // returns mailid in a string form and email in string form 
	
	// Cancels Mail
	MailID, err := strconv.Atoi(MailId) // Changes MailId from string to int 
	if err != nil {
		log.Fatal(err)
	}
	response, _ := client.DeleteMail(emailkey, MailID)
	fmt.Printf("%v\n", response)

	// Get Balance, Takes User's Key as the only arguement
	Balance := client.GetBalance(emailkey)
	fmt.Printf("%v\n", Balance)

	Value := client.GetLetter(emailkey, MailID) 
	fmt.Printf("Value: %v", Value) 
}
```
