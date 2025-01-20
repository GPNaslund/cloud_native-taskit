package dto

// Represents a task
type Task struct {
	Id      string `json:"id"`
	Title   string `json:"title" binding:"required"`
	Details string `json:"details" binding:"required"`
	IsDone  bool   `json:"is_done" binding:"required"`
}
