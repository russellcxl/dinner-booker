package http_

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type DinnerBot struct {
	token  string
	host   *url.URL
	client *http.Client
}

// NewDinnerBot should be initialised with host = "https://dinner.sea.com/"
func NewDinnerBot(host, token string) (*DinnerBot, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	bot := DinnerBot{
		host:   u,
		token:  token,
		client: &http.Client{},
	}
	return &bot, nil
}

// Get -- httpQuery currently takes in only TimeRange
func (bot *DinnerBot) Get(path string, httpQuery interface{}) (*http.Response, error) {
	u, err := url.Parse(bot.host.String())
	if err != nil {
		return nil, err
	}
	if httpQuery != nil {
		q, err := query.Values(httpQuery)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}
	u.Path = path

	log.Printf("accessing %s\n", u.String())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", bot.token))
	return bot.client.Do(req)
}

// Post -- used for ordering: https://dinner.sea.com/api/order/{menuID}; body uses form-data
func (bot *DinnerBot) Post(path string, body interface{}) (*http.Response, error) {
	u, err := url.Parse(bot.host.String())
	if err != nil {
		return nil, err
	}
	u.Path = path
	log.Printf("accessing %s\n", u.String())

	v, _ := query.Values(body)
	encoded := v.Encode()

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(encoded))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", bot.token))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // apparently, /api/order only accepts body as form-data, not JSON
	return bot.client.Do(req)
}