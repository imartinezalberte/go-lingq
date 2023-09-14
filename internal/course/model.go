package course

type (
	AddCourse struct {
		Title       string `example:"eventyr for barn"`
		Language    string `example:"no"`
		Description string `example:"Du vil lære å lese eventyr for barn"`
		Level       uint   `example:"2"`
		SourceURL   string `example:"https://www.barneforlaget.no/hør-så-mye-du-vil"`
		Tags        Tags
	}

	Tags []Tag
	Tag  string

	Course struct {
		ID               uint   `json:"id"`
		URL              string `json:"url"`
		Title            string `json:"title"`
		Description      string `json:"description"`
		ImageURL         string `json:"imageUrl"`
		OriginalImageURL string `json:"originalImageUrl"`
		LessonsCount     uint   `json:"lessonsCount"`
		Difficulty       uint   `json:"difficulty"`
		SharedByID       string `json:"sharedById"`
		SharedByName     string `json:"sharedByName"`
		SharedByImageURL string `json:"sharedByImageUrl"`
		Tags             Tags   `json:"tags"`
		Type             string `json:"type"`
	}
)
