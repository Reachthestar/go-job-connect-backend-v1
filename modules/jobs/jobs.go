package jobs

import "time"

type Job struct {
	ID            int64     `json:"id"`
    Description   string    `json:"description"`
    Qualification string    `json:"qualification"`
    Title         string    `json:"title"`
    Position      *string   `json:"position"`
    Type          string   	`json:"type"`
    Province      *string   `json:"province"`
    City          *string   `json:"city"`
    SalaryMin     *string   `json:"salaryMin"`
    SalaryMax     *string   `json:"salaryMax"`
    CreatedAt     time.Time `json:"createdAt"`
    UpdatedAt     time.Time `json:"updatedAt"`
    UserID        int64       `json:"userId"`
}