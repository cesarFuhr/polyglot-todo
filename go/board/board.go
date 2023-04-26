package board

import (
	"errors"
	"time"

	"github.com/cesarFuhr/go/polyglot-todo/task"
)

var (
	ErrEmptyName = errors.New("name should not be empty")
)

type Board struct {
	Name      string
	Tasks     []task.Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name string) (Board, error) {
	if name == "" {
		return Board{}, ErrEmptyName
	}

	b := Board{
		Name:      name,
		Tasks:     []task.Task{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return b, nil
}

func (b *Board) AddTask(t task.Task) {
	b.Tasks = append(b.Tasks, t)
}
