package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Define a struct to match the expected JSON response
type GhostResponse struct {
	Posts  []Post     `json:"posts,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

type Post struct {
	Title string `json:"title"`
	// Add other fields as necessary
}

type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	// Add other fields as necessary
}

func main() {
	fmt.Println("\n📝 Ghost Export to Markdown - Starting up...")

	if os.Getenv("GHOST_API_KEY") != "" {
		fmt.Println("🔑 `GHOST_API_KEY` detected, using:", os.Getenv("GHOST_API_KEY"))
	}

	if len(os.Args) < 3 {
		fmt.Println("Usage: ghost-export-to-markdown --url=<GHOST_URL> --api-key=<GHOST_API_KEY>")
		os.Exit(1)
	}

	apiUrl := flag.String("url", os.Getenv("GHOST_URL"), "Ghost URL")
	apiKey := flag.String("api-key", os.Getenv("GHOST_API_KEY"), "Ghost Admin API Key")

	flag.Parse()

	if *apiUrl == "" || *apiKey == "" {
		fmt.Println("Usage: ghost-export-to-markdown --url=<GHOST_URL> --api-key=<GHOST_API_KEY>")
		os.Exit(1)
	}

	fetchFromGhostAPI(*apiUrl, *apiKey)
	os.Exit(0)
}

func fetchFromGhostAPI(apiURL, apiKey string) {
	url := fmt.Sprintf("https://%s/ghost/api/v3/content/posts/?key=%s", apiURL, apiKey)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching from Ghost API:", err)
		return
	}

	defer resp.Body.Close()

	var result GhostResponse // Declare the result variable

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check for errors in the response
	if len(result.Errors) > 0 {
		for _, apiErr := range result.Errors {
			fmt.Printf("Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
		}
		return
	}

	fmt.Println("Response from Ghost API:", result.Posts) // Print the posts
}
