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
	Name      string      `json:"name"`
	CreatedAt int64       `json:"created_at"`
	UpdatedAt int64       `json:"updated_at"`
	Tasks     []task.Task `json:"tasks"`
}

func New(name string) (Board, error) {
	if name == "" {
		return Board{}, ErrEmptyName
	}

	b := Board{
		Name:      name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
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
	b.UpdatedAt = time.Now().Unix()
}

func (b *Board) UpdateTask(position int, t task.Task) error {
	if len(b.Tasks) == 0 || position >= len(b.Tasks) {
		return ErrInvalidPosition
	}

	b.UpdatedAt = time.Now().Unix()
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
	b.UpdatedAt = time.Now().Unix()
}
