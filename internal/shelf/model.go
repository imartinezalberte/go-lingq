package shelf

type (
	Shelves []Shelf
	Shelf   struct {
		ID    uint   `json:"id"`
		Code  string `json:"code"`
		Title string `json:"title"`
		Tabs  Tabs
	}

	Tabs []Tab
	Tab  struct {
		ApiURL   string `json:"apiUrl"`
		Selected bool   `json:"selected"`
		Level    uint   `json:"level"`
		Display  string `json:"display"`
		Title    string `json:"title"`
	}
)
