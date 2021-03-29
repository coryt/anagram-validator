package ports

// ValidateAnagramRequest defines model for checking anagram requests (POST /anagram/check)
type ValidateAnagramRequest struct {
	FirstWord  string `json:"firstWord"`
	SecondWord string `json:"secondWord"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}
