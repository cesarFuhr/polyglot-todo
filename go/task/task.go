package task

import (
	"errors"
	"time"
)

var (
	ErrEmptyTitle = errors.New("title should not be empty")
)

type Task struct {
	Done      bool
	Title     string
	CreatedAt time.Time
	DoneAt    time.Time
}

func New(title string) (Task, error) {
	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	return Task{
		Done:      false,
		Title:     title,
		CreatedAt: time.Now(),
		DoneAt:    time.Time{},
	}, nil
}
