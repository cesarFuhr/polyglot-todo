package main

import (
	"fmt"

	"github.com/cesarFuhr/go/polyglot-todo/board"
	"github.com/cesarFuhr/go/polyglot-todo/task"
)

func Add(b *board.Board, title string) error {
	t, err := task.New(title)
	if err != nil {
		return fmt.Errorf("creating a new task %w", err)
	}

	b.InsertTask(0, t)

	return nil
}

func List(b *board.Board) {
	fmt.Println(" ", b.Name)
	for i, t := range b.Tasks {
		fmt.Print(i+1, " ")
		if t.Done {
			fmt.Print("☑ ")
		} else {
			fmt.Print("☐ ")
		}
		fmt.Println(t.Title)
	}
}

func Done(b *board.Board, pos int) error {
	pos = pos - 1
	t, err := b.GetTask(pos)
	if err != nil {
		return fmt.Errorf("fetching task: %w", err)
	}

	t.Done = !t.Done
	err = b.UpdateTask(pos, t)
	if err != nil {
		return fmt.Errorf("updating task: %w", err)
	}

	return nil
}
