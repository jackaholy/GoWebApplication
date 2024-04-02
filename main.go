/*
Author: Jack Holy
Description: A Go program which generates a random quote from a famous individual. The quotes are taken from
the Quotable API which can be found at https://docs.quotable.io/. The user can click a button to generate a new
quote as many times as they want.
*/

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// Quote struct to represent a quote fetched from the API.
type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

/*
Get a random quote from the Quotable API.
@return Quote
*/
func fetchQuote() Quote {
	// Make a GET request to the "Quotable API"
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return Quote{}
	}

	// I got help from chat.openai.com for the following lines.
	// Decode the JSON response into a Quote struct.
	var quote Quote
	err = json.NewDecoder(resp.Body).Decode(&quote)
	if err != nil {
		return Quote{}
	}
	return quote
}

/*
Print the quote on the HTML page. Also prints errors if proper files do not exist.
I got help from chat.openai.com for the following line.
*/
func printQuote(w http.ResponseWriter, _ *http.Request) {
	// Get the quote
	quote := fetchQuote()

	// Display quote via index.html
	tmpl, err := template.ParseFiles("templates/index.html")
	// Check to see if the file exists, display error message otherwise.
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
	}

	// Render the template
	err = tmpl.Execute(w, quote)
	// Check to see if the template exists, display error message otherwise.
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

/*
Runs the web application through the localhost.
*/
func main() {
	http.HandleFunc("/", printQuote)

	// Set the port to run the local server on.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:")
	}
}
