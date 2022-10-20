package kopeechka

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (g *EmailClient) BuyEmail(ClientKey string, Domain string, Host string) (string, string) {
	g.ClientKey = ClientKey
	g.Domain = Domain
	g.Site = Host

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.kopeechka.store/mailbox-get-email?api=2.0&spa=1&site=%v&sender=twitch&regex=&mail_type=%v&token=%v", g.Site, g.Domain, g.ClientKey), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Status EmailClient
	err = json.Unmarshal(body, &Status)
	if err != nil {
		log.Fatal(err)
	}
	if Status.Status == "OK" {
		return Status.EmailId, Status.Email
	} else {
		return "NOT_OKAY", "ERROR"
	}
}

func (g *EmailClient) DeleteMail(ClientKey string, OrderId int) (string, error) {
	g.ClientKey = ClientKey
	g.OrderId = OrderId

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.kopeechka.store/mailbox-cancel?id=%v&token=%v", g.OrderId, g.ClientKey), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		return "OK", nil
	} else {
		return "NOT_OKAY", nil
	}
}

// func (g *EmailClient) GetEmail(ClientKey string, MailId int) (string, string) {
// 	g.ClientKey = ClientKey
// 	g.OrderId = MailId
// }

func (g *EmailClient) GetBalance(ClientKey string) float32 {
	g.ClientKey = ClientKey

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.kopeechka.store/user-balance?token=%v&type=json&api=2.0", g.ClientKey), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Client EmailClient
	err = json.Unmarshal(body, &Client)
	if err != nil {
		log.Fatal(err)
	}
	if Client.Balance > 0.01 {
		return Client.Balance
	} else {
		return 0
	}
}

func (g *EmailClient) GetLetter(ClientKey string, OrderId int) string {
	g.ClientKey = ClientKey
	g.OrderId = OrderId

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://api.kopeechka.store/mailbox-get-message?full=1&id=%v&token=%v&type=JSON&api=2.0", g.OrderId, g.ClientKey), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Client EmailClient
	err = json.Unmarshal(body, &Client)
	if err != nil {
		log.Fatal(err)
	}

	return Client.Value
}
