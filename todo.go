package gqlgen_todos

// Todo _
type Todo struct {
	ID     string               `json:"id"`
	Text   MyCustomStringScalar `json:"text"`
	Done   bool                 `json:"done"`
	UserID string               `json:"userId"`
}
