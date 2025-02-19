package ghost

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
