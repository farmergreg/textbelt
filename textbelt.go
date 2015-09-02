package textbelt

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// Send text messages to the USA
	TextbeltAPIusa = "http://textbelt.com/text"

	// Send text messages to Canada
	TextbeltAPIcanada = "http://textbelt.com/canada"

	// Send text messages internationally
	TextbeltAPIinternational = "http://textbelt.com/intl"
)

type textbeltResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Client is your entry point into the rest of the library.
type Client struct {
	apiURL string
}

// New creates a new client that will use the USA version
// of http://textbelt.com to send text messages.
func New() *Client {
	return NewClientFromURL(TextbeltAPIusa)
}

// NewClientFromURL creates a new client using a url of your choosing.
func NewClientFromURL(apiURL string) *Client {
	return &Client{
		apiURL: apiURL,
	}
}

// Text sends a text message.
func (c *Client) Text(phoneNumber string, message string) error {
	data := url.Values{
		"number":  {phoneNumber},
		"message": {message},
	}
	response, err := http.PostForm(c.apiURL, data)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseText, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	i := textbeltResponse{}
	err = json.Unmarshal([]byte(responseText), &i)

	if !i.Success {
		return errors.New(i.Message)
	}
	return nil
}
