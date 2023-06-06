package task

import (
	"errors"
	"time"
)

var (
	ErrEmptyTitle = errors.New("title should not be empty")
)

type Task struct {
	Done      bool   `json:"done"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at"`
	DoneAt    int64  `json:"done_at"`
}

func New(title string) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	t := Task{
		Done:      false,
		Title:     title,
		CreatedAt: time.Now().Unix(),
		DoneAt:    time.Time{}.Unix(),
	}
	return t, nil
}
