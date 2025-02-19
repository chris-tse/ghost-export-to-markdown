package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ghost-export-to-markdown/ghost"
	"ghost-export-to-markdown/markdown"

	"github.com/cqroot/prompt"
)

func main() {
	fmt.Print("\n📝 Ghost Export to Markdown - Starting up...\n\n")

	apiURL := flag.String("url", os.Getenv("GHOST_URL"), "Ghost URL")
	apiKey := flag.String("api-key", os.Getenv("GHOST_API_KEY"), "Ghost Admin API Key")
	dir := flag.String("dir", "ghost-export", "Directory to save the posts")

	flag.Parse()

	if *apiURL == "" {
		val, err := prompt.New().Ask("🌐 No URL detected. Enter your Ghost URL").Input(
			"",
		)

		if err != nil {
			log.Fatalln("Error during Ghost URL input:", err)
		}

		apiURL = &val
	}
	fmt.Println("🌐 Using Ghost URL: ", *apiURL)

	if *apiKey == "" {
		val, err := prompt.New().Ask("🔑 No API key detected. Enter your Ghost API key").Input(
			"",
		)

		if err != nil {
			log.Fatalln("Error during Ghost API key input:", err)
		}

		apiKey = &val
	}
	fmt.Printf("🔑 Using Ghost API Key: %s\n\n", *apiKey)

	if *apiURL == "" || *apiKey == "" {
		fmt.Println("Usage: ghost-export-to-markdown --url=your.ghost.url --api-key=Your1Ghost2Api3Key")
		fmt.Println("You can also set GHOST_URL and GHOST_API_KEY as environment variables")
		os.Exit(1)
	}

	if err := os.MkdirAll(*dir, 0755); err != nil {
		log.Fatalln("Error creating directory:", err)
	}

	fmt.Println("🔍 Fetching posts from Ghost API...")

	client := ghost.NewClient(*apiURL, *apiKey)

	posts, err := client.FetchPosts()
	if err != nil {
		log.Fatalln("Error during Ghost API fetch:", err)
	}

	for _, post := range posts {
		fmt.Println("💾 Saving post:", post.Title)

		convertedPost, err := markdown.ConvertPost(post)
		if err != nil {
			fmt.Printf("Error during post conversion: %s\n", err)
		}

		os.WriteFile(*dir+"/"+post.Slug+".md", convertedPost, 0644)
	}

	os.Exit(0)
}

// func savePostToMarkdown(post Post, dir string) error {
// 	// Convert HTML to Markdown
// 	markdown, err := htmltomarkdown.ConvertString(post.Html)

// 	if err != nil {
// 		fmt.Printf("Error converting HTML to Markdown: %s\n", post.Slug)
// 		return err
// 	}

// 	// Prepare frontmatter
// 	frontmatter := fmt.Sprintf("---\nTitle: %s\nPublished: %s\nTags: %v\n---\n", post.Title, post.PublishedAt, post.Tags)

// 	// Combine frontmatter and markdown content
// 	finalContent := frontmatter + markdown

// 	// Save to a markdown file
// 	fileName := fmt.Sprintf("%s/%s.md", dir, post.Slug)

// 	err = os.WriteFile(fileName, []byte(finalContent), 0644)

// 	if err != nil {
// 		fmt.Printf("Error saving post %s to file\n: %v", post.Slug, err)
// 		return err
// 	}
// }
