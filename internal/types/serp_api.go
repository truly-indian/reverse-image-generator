package types

type SearchMetadata struct {
	ID               string  `json:"id"`
	Status           string  `json:"status"`
	JSONEndpoint     string  `json:"json_endpoint"`
	CreatedAt        string  `json:"created_at"`
	ProcessedAt      string  `json:"processed_at"`
	GoogleLensURL    string  `json:"google_lens_url"`
	RawHTMLFile      string  `json:"raw_html_file"`
	PrettifyHTMLFile string  `json:"prettify_html_file"`
	TotalTimeTaken   float64 `json:"total_time_taken"`
}

// SearchParameters contains parameters used in the search
type SearchParameters struct {
	Engine string `json:"engine"`
	URL    string `json:"url"`
}

// VisualMatch represents a visual match result
type VisualMatch struct {
	Position   int    `json:"position"`
	Title      string `json:"title"`
	Link       string `json:"link"`
	Source     string `json:"source"`
	SourceIcon string `json:"source_icon"`
	Thumbnail  string `json:"thumbnail"`
}

type ImageSourcesSearch struct {
	PageToken   string `json:"page_token"`
	SerpapiLink string `json:"serpapi_link"`
}

type SerpAPIResponse struct {
	SearchMetadata     SearchMetadata     `json:"search_metadata"`
	SearchParameters   SearchParameters   `json:"search_parameters"`
	VisualMatches      []VisualMatch      `json:"visual_matches"`
	ImageSourcesSearch ImageSourcesSearch `json:"image_sources_search"`
}
