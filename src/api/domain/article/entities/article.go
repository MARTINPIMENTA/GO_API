package entities

// Entities for article functionalities.
type (
	Article struct {
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
	}

	Articles []Article
)
