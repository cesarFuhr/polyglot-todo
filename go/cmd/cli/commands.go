package main

import (
	"fmt"
	"time"

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
	t.DoneAt = time.Now().Unix()
	err = b.UpdateTask(pos, t)
	if err != nil {
		return fmt.Errorf("updating task: %w", err)
	}

	return nil
}

func Delete(b *board.Board, pos int) error {
	pos = pos - 1
	_, err := b.GetTask(pos)
	if err != nil {
		return fmt.Errorf("fetching task: %w", err)
	}

	b.RemoveTask(pos)

	return nil
}

func Update(b *board.Board, pos int, title string) error {
	pos = pos - 1
	t, err := b.GetTask(pos)
	if err != nil {
		return fmt.Errorf("fetching task: %w", err)
	}

	t.Title = title
	err = b.UpdateTask(pos, t)
	if err != nil {
		return fmt.Errorf("updating task: %w", err)
	}

	return nil
}
