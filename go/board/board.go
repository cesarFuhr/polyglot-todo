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
	CreatedAt time.Time
	UpdatedAt time.Time
	Tasks     []task.Task
}

func New(name string) (Board, error) {
	if name == "" {
		return Board{}, ErrEmptyName
	}

	b := Board{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tasks:     []task.Task{},
	}
	return b, nil
}

func (b *Board) GetTasks() []task.Task {
	return b.Tasks
}

func (b *Board) GetTask(position int) (task.Task, error) {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		return task.Task{}, ErrInvalidPosition
	}

	return b.Tasks[position], nil
}

func (b *Board) InsertTask(position int, t task.Task) {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		b.Tasks = append(b.Tasks, t)
		return
	}

	left := b.Tasks[:position]
	right := b.Tasks[position:]

	result := make([]task.Task, 0, len(b.Tasks)+1)
	result = append(result, left...)
	result = append(result, t)
	result = append(result, right...)

	b.Tasks = result
	b.UpdatedAt = time.Now()
}

func (b *Board) UpdateTask(position int, t task.Task) error {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		return ErrInvalidPosition
	}

	b.UpdatedAt = time.Now()
	b.Tasks[position] = t

	return nil
}

func (b *Board) RemoveTask(position int) {
	if position >= len(b.Tasks) {
		return
	}
	if len(b.Tasks) == 0 || position == len(b.Tasks)-1 {
		b.Tasks = b.Tasks[:position]
		return
	}

	left := b.Tasks[:position]
	right := b.Tasks[position+1:]
	b.Tasks = append(left, right...)
	b.UpdatedAt = time.Now()
}
