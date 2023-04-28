package main

import "strings"

type CommandTask struct {
	Name string
}

func NewCommandTask(args []string) CommandTask {
	name := strings.Join(args[1:], " ")

	return CommandTask{
		Name: name,
	}
}

func (c *CommandTask) AddTask() error {
	// create Todo board
	// add task
	// save the board

	return nil
}
