package todo

import "time"

type Status string

const (
	IN_PROGRESS Status = "IN_PROGRESS"
	COMPLETED   Status = "COMPLETED"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Status      Status `json:"status"`
}
