package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

type GhostResponse struct {
	Posts  []Post     `json:"posts,omitempty"`
	Errors []APIError `json:"errors,omitempty"`
}

type Post struct {
	Title               string `json:"title"`
	Html                string `json:"html"`
	Excerpt             string `json:"excerpt"`
	ReadingTime         int    `json:"reading_time"`
	PublishedAt         string `json:"published_at"`
	FeatureImageSrc     string `json:"feature_image"`
	FeatureImageAlt     string `json:"feature_image_alt"`
	FeatureImageCaption string `json:"feature_image_caption"`
	Slug                string `json:"slug"`
}

type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	// Add other fields as necessary
}

func main() {
	fmt.Print("\n📝 Ghost Export to Markdown - Starting up...\n\n")

	apiUrl := flag.String("url", os.Getenv("GHOST_URL"), "Ghost URL")
	apiKey := flag.String("api-key", os.Getenv("GHOST_API_KEY"), "Ghost Admin API Key")
	dir := flag.String("dir", "ghost-export", "Directory to save the posts")

	flag.Parse()

	if *apiUrl == "" || *apiKey == "" {
		fmt.Println("Usage: ghost-export-to-markdown --url=your.ghost.url --api-key=Your1Ghost2Api3Key")
		fmt.Println("You can also set GHOST_URL and GHOST_API_KEY as environment variables")
		os.Exit(1)
	}

	fmt.Println("🌐 Using Ghost URL: ", *apiUrl)
	fmt.Printf("🔑 Using Ghost API Key: %s\n\n", *apiKey)

	fmt.Println("🔍 Fetching posts from Ghost API...")
	posts, err := fetchFromGhostAPI(*apiUrl, *apiKey)

	if err != nil {
		log.Fatalln("Error during Ghost API fetch:", err)
	}

	if err := os.MkdirAll(*dir, 0755); err != nil {
		log.Fatalln("Error creating directory:", err)
	}

	for _, post := range posts {
		savePostToMarkdown(post, *dir)
	}

	os.Exit(0)
}

func fetchFromGhostAPI(apiURL, apiKey string) ([]Post, error) {
	url := fmt.Sprintf("https://%s/ghost/api/content/posts/?key=%s", apiURL, apiKey)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result GhostResponse // Declare the result variable

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Check for errors in the response
	if len(result.Errors) > 0 {
		for _, apiErr := range result.Errors {
			fmt.Printf("Error: %s (Code: %s)\n", apiErr.Message, apiErr.Code)
		}
		return nil, err
	}

	return result.Posts, nil
}

func savePostToMarkdown(post Post, dir string) {
	// Convert HTML to Markdown
	markdown, err := htmltomarkdown.ConvertString(post.Html)

	if err != nil {
		fmt.Printf("Error converting HTML to Markdown: %s\n", post.Slug)
		return
	}

	// Prepare frontmatter
	frontmatter := fmt.Sprintf("---\nTitle: %s\nPublished: %s\nTags: %v\n---\n", post.Title, post.PublishedAt, post.Tags)

	// Combine frontmatter and markdown content
	finalContent := frontmatter + markdown

	// Save to a markdown file
	fileName := fmt.Sprintf("%s/%s.md", dir, post.Slug)

	err := os.WriteFile(fileName, []byte(finalContent), 0644)

	if err != nil {
		fmt.Printf("Error saving post %s to file\n: %v", post.Slug, err)
		return
	}
}
