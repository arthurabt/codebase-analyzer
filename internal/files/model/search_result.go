package model

// FileSearchResult represents all the reults of a code pattern search in a file
type FileSearchResult struct {
	File    string
	FileURL string
	Error   error
	Results []SearchResult
}

// SearchResult represents the result of a code pattern search in a file
type SearchResult struct {
	LineNumber       int
	CodeLines        string
	CodeLineURL      string
	FlagKey          string
	FlagDefaultValue string
}
