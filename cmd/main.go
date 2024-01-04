package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// words represents a sample list of words to choose from
var words = []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince", "raspberry", "strawberry", "tangerine", "ugli", "vanilla", "watermelon", "xigua", "yam", "zucchini"}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the REST server!")
	})

	// Define the handler for the random words endpoint
	http.HandleFunc("/random-words", randomWordsHandler)

	// Start the server on port 8080
	fmt.Println("Server is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// randomWordsHandler handles requests for random words
func randomWordsHandler(w http.ResponseWriter, r *http.Request) {
	selectedWords := getRandomWords(5)
	jsonResponse, err := json.Marshal(selectedWords)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// getRandomWords selects n random words from the words slice
func getRandomWords(n int) []string {
	var selected []string
	for i := 0; i < n; i++ {
		selected = append(selected, words[rand.Intn(len(words))])
	}
	return selected
}
