package alerts

import (
	"errors"
	"time"
)

// Alert represent an alert
type Alert struct {
	ID       string `json:"id"` // Should be an UUID
	UserID   string `json:"user_id"`
	Date     string `json:"date"` // Date with format YYYY-MM-DD (ISO-8601)
	Name     string `json:"name"`
	MinPrice int    `json:"min_price,omitempty"`
	MaxPrice int    `json:"max_price,omitempty"`
	City     string `json:"city"`
}

// ErrInvalidPriceRange is returned when the min price is higher than the max price
var ErrInvalidPriceRange = errors.New("Max price is lower than min price")

// NewAlert create an alert
func NewAlert(id, userID, date, name string, minP, maxP int, city string) (Alert, error) {
	// Validate the date
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Alert{}, err
	}

	// Validate price range
	if maxP < minP {
		return Alert{}, ErrInvalidPriceRange
	}

	return Alert{
		ID:       id,
		UserID:   userID,
		Date:     date,
		Name:     name,
		MaxPrice: maxP,
		MinPrice: minP,
		City:     city,
	}, nil
}
