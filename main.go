package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ItsMerkz/Kopeechka_Wrapper/kopeechka"
)

func main() {
	// Email Client!
	var emailkey string = "" // Users Key

	client := new(kopeechka.EmailClient)

	// Buys Email Taking 3 arguements, returning mailid and email
	MailId, Email := client.BuyEmail(emailkey, "mail.ru", "twitch.tv")
	fmt.Printf("%v | %v\n", MailId, Email)

	// Cancels Mail
	MailID, err := strconv.Atoi(MailId) // changes MailId from BuyEmail function to an int from a string
	if err != nil {
		log.Fatal(err)
	}
	response, _ := client.DeleteMail(emailkey, MailID)
	fmt.Printf("%v\n", response)

	// Get Balance, Takes User's Key as the only arguement
	Balance := client.GetBalance(emailkey)
	fmt.Printf("%v\n", Balance)

  // Get Letter 
	Value := client.GetLetter(emailkey, 1441288813)
	fmt.Printf("%v\n", Value)
}
