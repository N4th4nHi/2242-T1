package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Bio represents the personal information of an individual
type Bio struct {
	Name        string
	Age         int
	Occupation  string
	Description string
}

// QOTD represents a quote of the day
type QOTD struct {
	Quote string
}

// Greeting represents a greeting message
type Greeting struct {
	Time      string
	DayOfWeek string
}

func main() {
	// Handle the root URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a Bio instance
		bio := Bio{
			Name:        "Nathan Hislop",
			Age:         19,
			Occupation:  "Student",
			Description: "I'm a student who is currently about to finish school...at last",
		}

		// Parse the biography.html template
		tmpl, err := template.ParseFiles("biography.html")
		if err != nil {
			// If there is an error, return a server error to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the biography template with the Bio instance and write the output to the response writer
		err = tmpl.Execute(w, bio)
		if err != nil {
			// If there is an error, return a server error to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Create a list of quotes
	quotes := []string{
		"I actually don’t like thinking. I think people think I like to think a lot. And I don’t. I do not like to think at all. - Kanye West",
		"I’ve never really wanted to go to Japan. Simply because I don’t like eating fish. And I know that’s very popular out there in Africa - Britney Spears",
		"No, no, I didn’t go to England, I went to London. - Jaden Smith",
		"Smoking kills. If you’re killed, you’ve lost an important part of your life - Brooke Shields",
		"The only happy artist is a dead artist, because only then you can’t change. After I die, I’ll probably come back as a paintbrush. - Sylvester Stallone",
	}

	// Handle the /quotes URL path
	http.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		// Set the random seed
		rand.Seed(time.Now().Unix())

		// Get a random index in the quotes list
		index := rand.Intn(len(quotes))

		// Create a QOTD instance with the selected quote
		quote := quotes[index]
		qotd := QOTD{
			Quote: quote,
		}

		// Parse the quotes.html template
		tmpl, err := template.ParseFiles("quotes.html")
		if err != nil {
			// If there is an error, return a server error to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the quotes template with the QOTD instance and write the output to the response writer
		err = tmpl.Execute(w, qotd)
		if err != nil {
			// If there is an error, return a server error to the client
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Handle the /greeting URL

	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		// Get the current time
		now := time.Now()
		// Get the current day of the week as a string
		dayOfWeek := now.Weekday().String()

		// Create a new Greeting object with the current time and day of the week
		greeting := Greeting{
			Time:      now.Format(time.Kitchen),
			DayOfWeek: dayOfWeek,
		}

		// Parse the greeting.html template file
		tmpl, err := template.ParseFiles("greeting.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Render the template with the Greeting object data
		err = tmpl.Execute(w, greeting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
