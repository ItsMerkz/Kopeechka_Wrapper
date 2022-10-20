# Kopeechka Wrapper (Golang)

# Functions:
* Get Balance
* Buy Email 
* Cancel Email 
* Get Letter 

### Full Docuemtation Of Client Below!

**Buy Email**
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

```
