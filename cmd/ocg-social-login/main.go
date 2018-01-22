package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aweisser/ocg-poc/io/githubaccount"
	"github.com/danilopolani/gocialite"
)

var gocial = gocialite.NewDispatcher()

// Call with two env params GITHUB_CLIENT_ID and GITHUB_CLIENT_SECRET.
// Create Client-ID and Secret via GitHub OAuth Apps in Developer Settings
// GITHUB_CLIENT_ID="xxxxxxxxxxxxx" GITHUB_CLIENT_SECRET="xxxxxxxxxxxxxxxxxxxxxxxxxx" go run main.go
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/auth/github", redirectHandler)
	http.HandleFunc("/auth/github/callback", callbackHandler)
	http.ListenAndServe(":9090", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Social Logins"))
}

// Redirect to correct oAuth URL
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	authURL, err := gocial.New().
		Driver("github").                 // Set provider
		Scopes([]string{"repo", "user"}). // Set optional scope(s)
		Redirect(                         //
			clientID,                                     // Client ID
			clientSecret,                                 // Client Secret
			"http://localhost:9090/auth/github/callback", // Redirect URL
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	http.Redirect(w, r, authURL, http.StatusFound) // Redirect with 302 HTTP code
}

// Redirect to correct oAuth URL
// Handle callback of provider
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve query params for code and state
	codeParam, _ := r.URL.Query()["code"]
	code := codeParam[0]

	stateParam, _ := r.URL.Query()["state"]
	state := stateParam[0]

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

	// Use the AccessToken to call the GitHub REST and GraphQL API
	client, err2 := githubaccount.NewClient(token.AccessToken)
	if err2 != nil {
		w.Write([]byte("Error: " + err2.Error()))
		return
	}
	stats, err3 := githubaccount.StatsOf(client)
	if err3 != nil {
		w.Write([]byte("Error: " + err3.Error()))
		return
	}

	// If no errors, show provider name and stats
	w.Write([]byte("Hi, " + user.FullName + "\n"))
	w.Write([]byte("Stats: " + *stats.Login + "\n"))
}
