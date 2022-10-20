package kopeechka

import "net/http"

type EmailClient struct {
	Client    *http.Client
	ClientKey string
	Domain    string
	Site      string
	OrderId   int
	Status    string  `json:"status"`
	EmailId   string  `json:"id"`
	Email     string  `json:"mail"`
	Balance   float32 `json:"balance"`
	Value     string  `json:"value"`
}
