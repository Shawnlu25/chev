package keep

// ExportedNote is the json format model of Google Keep Export
type ExportedNote struct {
	Color                   string   `json:"color"`
	IsTrashed               bool     `json:"isTrashed"`
	IsPinned                bool     `json:"isPinned"`
	IsArchived              bool     `json:"isArchived"`
	TextContent             string   `json:"textContent"`
	Title                   string   `json:"title"`
	UserEditedTimestampUsec uint64   `json:"userEditedTimestampUsec"`
	Labels                  *[]Label `json:"labels"`
}

// Label is how Google Keep models a single label tagged to a note
type Label struct {
	Name string `json:"name"`
}
