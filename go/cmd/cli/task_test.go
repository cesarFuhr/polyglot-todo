package main

import "testing"

func TestAddTask(t *testing.T) {
	t.Run("should create a task in the default todo board", func(t *testing.T) {
		args := []string{"task", "Name", "of", "the", "task"}

		cmd := NewCommandTask(args)
		cmd.AddTask()

	})
}
