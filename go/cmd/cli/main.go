package main

import (
	"flag"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const (
	CmdTask = "task"
)

func run() error {
	flag.Parse()

	// Branch to the requested command.
	switch flag.Arg(0) {
	case CmdTask:
		c := NewCommandTask(flag.Args())

	default:
		flag.Usage()
	}

	return nil
}

// Create a task.
// ./todo task New task name
