package models

type Event struct {
	OrderType   string `json:"orderType"`
	SessionID   string `json:"sessionId"`
	Card        string `json:"card"`
	EventDate   string `json:"eventDate"`
	WebsiteUrl  string `json:"websiteUrl"`
	IsDelivered bool   `json:"-"`
}
