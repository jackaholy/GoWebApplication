package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Quote struct to represent a quote fetched from the API
type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func fetchQuote() Quote {
	// Make a GET request to the "Quotable API"
	resp, err := http.Get("https://api.quotable.io/random")

	func(Body io.ReadCloser) {
	}(resp.Body)

	// Decode the JSON response into a Quote struct
	var quote Quote
	err = json.NewDecoder(resp.Body).Decode(&quote)

	// Check if quote exists
	if err != nil {
		return Quote{}
	}

	return quote
}

func printQuote(w http.ResponseWriter, _ *http.Request) {
	// Fetch a random quote
	quote := fetchQuote()

	// Format the quote to be printed
	fmt.Fprintf(w, "Random Quote:\n%s\n- %s", quote.Content, quote.Author)
}

func main() {
	http.HandleFunc("/", printQuote)
	// Set the port to run the local server on
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
