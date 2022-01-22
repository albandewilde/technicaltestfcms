package alerts

// Alert represent an alert
type Alert struct {
	ID       string `json:"id"` // Should be an UUID
	UserID   string `json:"user_id"`
	Date     string `json:"date"` // Date with format YYYY-MM-DD (ISO-8601)
	Name     string `json:"name"`
	MaxPrice int    `json:"max_price"`
	MinPrice int    `json:"min_price"`
	City     string `json:"city"`
}

// NewAlert create an alert
func NewAlert(id, userID, date, name string, maxP, minP int, city string) Alert {
	return Alert{
		ID:       id,
		UserID:   userID,
		Date:     date,
		Name:     name,
		MaxPrice: maxP,
		MinPrice: minP,
		City:     city,
	}
}
