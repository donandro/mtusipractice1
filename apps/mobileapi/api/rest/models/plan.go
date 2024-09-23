package models

type GetPlansResponse struct {
	Plans []Plan `json:"plans"`
}

type Plan struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ReminderTime string `json:"reminderTime"`
	Frequency    string `json:"frequency"`
	DaysOfWeek   string `json:"daysOfWeek"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Pill         Pill   `json:"pill"`
}
