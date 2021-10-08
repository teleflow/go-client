package teleflow

type PaginationItems struct {
	ItemsCount int64       `json:"items_count,omitempty"`
	PagesCount int         `json:"pages_count,omitempty"`
	NextLink   string      `json:"next_link,omitempty"`
	PrevLink   string      `json:"prev_link,omitempty"`
	FirstLink  string      `json:"first_link,omitempty"`
	LastLink   string      `json:"last_link,omitempty"`
	Page       int         `json:"page,omitempty"`
	Items      interface{} `json:"items"`
}
