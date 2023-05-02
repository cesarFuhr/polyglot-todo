package main

import (
	"errors"
	"flag"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var f flag.FlagSet
	f.Bool("l", false, "Lists all current tasks.")
	f.Bool("a", false, "Adds a new task to the board.")
	f.Int("d", 0, "Marks a task as done.")
	f.Int("D", 0, "Deletes a task.")
	f.Int("m", 0, "Moves a task to the given poisition.")
	f.Int("u", 0, "Updates the title of the given task.")
	f.Parse(flag.Args())

	if err := validateFlags(f); err != nil {
		return err
	}

	flags := newFlags(f)

	switch {
	case flags.list:
		// call some list function
	}

	return nil
}

func validateFlags(f flag.FlagSet) error {
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

func newFlags(f flag.FlagSet) flags {
	// If no flag was set consider as a list command.
	if f.NFlag() == 0 {
		return flags{list: true}
	}

	return flags{}
}
