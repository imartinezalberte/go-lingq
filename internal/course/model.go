package course

type (
	AddCourse struct {
		Title       string `json:"title"       example:"eventyr for barn"`
		Language    string `json:"language"    example:"no"`
		Description string `json:"description" example:"Du vil lære å lese eventyr for barn"`
		Level       uint   `json:"level"       example:"2"`
		SourceURL   string `json:"sourceUrl"   example:"https://www.barneforlaget.no/hør-så-mye-du-vil"`
		Tags        Tags   `json:"tags"`
	}

	Tags []Tag
	Tag  string

	Courses []Course
	Course  struct {
		ID               uint    `json:"id"`
		URL              string  `json:"url"`
		Title            string  `json:"title"`
		Description      string  `json:"description"`
		ImageURL         string  `json:"imageUrl"`
		OriginalImageURL string  `json:"originalImageUrl"`
		LessonsCount     uint    `json:"lessonsCount"`
		Difficulty       float32 `json:"difficulty"`
		SharedByID       string  `json:"sharedById"`
		SharedByName     string  `json:"sharedByName"`
		SharedByImageURL string  `json:"sharedByImageUrl"`
		Tags             Tags    `json:"tags"`
		Type             string  `json:"type"`
	}

	CourseID       string
	CounterCourses map[CourseID]CounterCourse
	CounterCourse  struct {
		CardsCount        uint    `json:"cardsCount"`
		HasFlags          bool    `json:"hasFlags"`
		KnownWordsCount   uint    `json:"knownWordsCount"`
		TotalWordsCount   uint    `json:"totalWordsCount"`
		RoseGiven         bool    `json:"roseGiven"`
		RosesCount        uint    `json:"rosesCount"`
		LessonsCount      uint    `json:"lessonsCount"`
		Difficulty        float32 `json:"difficulty"`
		IsTaken           bool    `json:"isTaken"`
		NewWordsCount     uint    `json:"newWordsCount"`
		CourseID          int     `json:"pk"`
		Progress          any     `json:"progress"`
		IsCompletelyTaken bool    `json:"isCompletelyTaken"`
		UniqueWordsCount  uint    `json:"uniqueWordsCount"`
	}
)
