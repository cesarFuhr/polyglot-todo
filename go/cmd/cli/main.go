package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"
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

	switch {
	case flags.add:
		AddTask(nil, strings.Join(f.Args(), " "))
	}

	return nil
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
