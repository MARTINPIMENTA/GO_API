package entities

// Entities for article functionalities.
type (
	Article struct {
		UserID  string `json:"user_id"`
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
	}

	Articles []Article
)
