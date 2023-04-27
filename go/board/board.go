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

	tasks []task.Task
}

func New(name string) (Board, error) {
	if name == "" {
		return Board{}, ErrEmptyName
	}

	b := Board{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		tasks:     []task.Task{},
	}
	return b, nil
}

func (b *Board) GetTasks() []task.Task {
	return b.tasks
}

func (b *Board) GetTask(position int) (task.Task, error) {
	if len(b.tasks) == 0 || position >= len(b.tasks) {
		return task.Task{}, ErrInvalidPosition
	}

	return b.tasks[position], nil
}

func (b *Board) InsertTask(position int, t task.Task) {
	if len(b.tasks) == 0 || position >= len(b.tasks) {
		b.tasks = append(b.tasks, t)
		return
	}

	left := b.tasks[:position]
	right := b.tasks[position:]

	temp := append(left, t)
	b.tasks = append(temp, right...)
	b.UpdatedAt = time.Now()
}

func (b *Board) UpdateTask(position int, t task.Task) error {
	if len(b.tasks) == 0 || position >= len(b.tasks) {
		return ErrInvalidPosition
	}

	b.UpdatedAt = time.Now()
	b.tasks[position] = t

	return nil
}

func (b *Board) RemoveTask(position int) {
	if len(b.tasks) == 0 || position >= len(b.tasks)-1 {
		b.tasks = b.tasks[:position]
		return
	}

	left := b.tasks[:position]
	right := b.tasks[position+1:]
	b.tasks = append(left, right...)
	b.UpdatedAt = time.Now()
}
