package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cesarFuhr/go/polyglot-todo/board"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.Bool("l", false, "Lists all current tasks.")
	f.Bool("a", false, "Adds a new task to the board.")
	f.Int("d", 0, "Marks a task as done.")
	f.Int("D", 0, "Deletes a task.")
	f.Int("m", 0, "Moves a task to the given poisition.")
	f.Int("u", 0, "Updates the title of the given task.")
	f.Parse(os.Args[1:])

	// Debugging: print all flags value
	f.VisitAll(func(f *flag.Flag) {
		log.Printf("%+v", f)
	})

	if err := validateFlags(f); err != nil {
		return err
	}

	flags := newFlags(f)
	log.Printf("%+v", flags)

	b, err := loadBoard(".todo.json")
	if err != nil {
		return err
	}

	switch {
	case flags.add:
		AddTask(b, strings.Join(f.Args(), " "))
	}

	return saveBoard(".todo.json", b)
}

func validateFlags(f *flag.FlagSet) error {
	if f.NFlag() > 1 {
		return errors.New("cannot use multiple flags")
	}

	return nil
}

type flags struct {
	list   bool
	add    bool
	done   int
	delete int
	move   int
	update int
}

func newFlags(f *flag.FlagSet) flags {
	// If no flag was set consider as a list command.
	if f.NFlag() == 0 {
		return flags{list: true}
	}

	add := f.Lookup("a")
	log.Printf("%+v", add)
	return flags{add: add.Value.String() == "true"}
}

func loadBoard(path string) (*board.Board, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b board.Board
	err = json.NewDecoder(file).Decode(&b)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("decoding board %w", err)
	}

	if b.Name == "" {
		b, err = board.New("TODO")
		if err != nil {
			return nil, fmt.Errorf("creating a board: %w", err)
		}
	}

	return &b, nil
}

func saveBoard(path string, b *board.Board) error {
	file, err := os.OpenFile(".todo.json", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(b)
	if err != nil {
		return fmt.Errorf("encoding board %w", err)
	}

	return nil
}
