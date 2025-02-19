package markdown

import (
	"fmt"
	"ghost-export-to-markdown/ghost"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func ConvertPost(post ghost.Post) ([]byte, error) {

	markdown, err := htmltomarkdown.ConvertString(post.Html)
	if err != nil {
		return nil, err
	}

	frontmatter := fmt.Sprintf(`---
title: %s
published_at: %s
reading_time: %d
excerpt: %s
slug: %s
---

`, post.Title, post.PublishedAt, post.ReadingTime, post.Excerpt, post.Slug)

	return []byte(frontmatter + markdown), nil
}
