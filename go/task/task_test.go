package task_test

import (
	"testing"

	"github.com/cesarFuhr/go/polyglot-todo/task"
)

func TestNewTask(t *testing.T) {
	t.Run("should give an error if the title is empty ", func(t *testing.T) {
		empty := ""

		expected := task.ErrEmptyTitle
		_, actual := task.New(empty)

		if expected != actual {
			t.Fatalf("expected %v, received %v", expected, actual)
		}
	})
}
