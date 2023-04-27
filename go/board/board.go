package board

import (
	"errors"
	"time"

	"github.com/cesarFuhr/go/polyglot-todo/task"
)

var (
	ErrEmptyName       = errors.New("name should not be empty")
	ErrInvalidPosition = errors.New("invalid board position")
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

func (b *Board) UpdateTask(position int, t task.Task) error {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		return ErrInvalidPosition
	}

	b.Tasks[position] = t

	return nil
}

func (b *Board) InsertTask(position int, t task.Task) {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		b.Tasks = append(b.Tasks, t)
		return
	}

	left := b.Tasks[:position]
	right := b.Tasks[position:]

	temp := append(left, t)
	b.Tasks = append(temp, right...)
}

func (b *Board) RemoveTask(position int) {
	if len(b.Tasks) == 0 || position >= len(b.Tasks)-1 {
		b.Tasks = b.Tasks[:position]
		return
	}

	left := b.Tasks[:position]
	right := b.Tasks[position+1:]
	b.Tasks = append(left, right...)
}
