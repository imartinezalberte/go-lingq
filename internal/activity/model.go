package activity

import "github.com/imartinezalberte/go-lingq/internal/entities"

type (
	Activity struct {
		ActivityApple  string        `json:"activityApple"`
		Coints         uint          `json:"coins"`
		Language       string        `json:"string"`
		DailyGoal      uint          `json:"dailyGoal"`
		StreakDays     uint          `json:"streakDays"`
		ActivityLevel  ActivityLevel `json:"activityLevel"`
		DailyMetric    string        `json:"dailyMetric"`
		DailyScores    DailyScores   `json:"dailyScores"`
		KnownWords     uint          `json:"knownWords"`
		IsStreakBroken bool          `json:"isStreakBroken"`
	}

	DailyScores []DailyScore
	DailyScore  struct {
		Date          entities.DateOnly `json:"date"`
		DayOfWeek     string            `json:"dayOfWeek"`
		ActivityLevel ActivityLevel     `json:"activityLevel"`
		Score         uint              `json:"score"`
	}

	ActivityLevel struct {
		Score uint `json:"score"`
		ID    uint `json:"id"`
	}

	Lingqs []Lingq
	Lingq  struct {
		Date          entities.DateOnly `json:"date"`
		DayOfWeek     string            `json:"dayOfWeek"`
		LingqsCreated uint              `json:"lingqsCreated"`
	}
)
