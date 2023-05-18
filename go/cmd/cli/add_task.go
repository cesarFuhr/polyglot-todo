package main

import (
	"fmt"

	"github.com/cesarFuhr/go/polyglot-todo/board"
	"github.com/cesarFuhr/go/polyglot-todo/task"
)

func AddTask(b *board.Board, title string) error {
	t, err := task.New(title)
	if err != nil {
		return fmt.Errorf("creating a new task %w", err)
	}

	b.InsertTask(0, t)

	return nil
}
