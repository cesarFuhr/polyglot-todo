package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/cesarFuhr/go/polyglot-todo/board"
	"github.com/cesarFuhr/go/polyglot-todo/task"
)

func AddTask(w io.ReadWriter, title string) error {
	var b board.Board
	err := json.NewDecoder(w).Decode(&b)
	if err != nil {
		return fmt.Errorf("decoding board %w", err)
	}

	// Hardcoded, maybe will implement multiple boards in
	// the future
	b.Name = "TODO"
	t, err := task.New(title)
	if err != nil {
		return fmt.Errorf("creating a new task %w", err)
	}

	b.InsertTask(0, t)
	log.Printf("%+v\n", b)

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		return fmt.Errorf("encoding board %w", err)
	}

	return nil
}
