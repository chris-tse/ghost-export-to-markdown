package ghost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Client handles all Ghost API interactions
type Client struct {
	apiURL string
	apiKey string
	http   *http.Client
}

// NewClient creates a new Ghost API client
func NewClient(apiURL, apiKey string) *Client {

	return &Client{
		apiURL: strings.TrimPrefix(apiURL, "https://"),
		apiKey: apiKey,
		http:   &http.Client{},
	}
}

// FetchPosts retrieves all posts from the Ghost API
func (c *Client) FetchPosts() ([]Post, error) {
	url := fmt.Sprintf("https://%s/ghost/api/content/posts/?key=%s", c.apiURL, c.apiKey)

	resp, err := c.http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Posts  []Post     `json:"posts,omitempty"`
		Errors []APIError `json:"errors,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("API error: %s (code: %s)",
			result.Errors[0].Message,
			result.Errors[0].Code,
		)
	}

	return result.Posts, nil
}
