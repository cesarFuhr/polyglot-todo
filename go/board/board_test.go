package board_test

import (
	"fmt"
	"testing"

	"github.com/cesarFuhr/go/polyglot-todo/board"
	"github.com/cesarFuhr/go/polyglot-todo/task"
)

func TestNewBoard(t *testing.T) {
	t.Run("should return a valid board", func(t *testing.T) {
		name := "board name"
		board, err := board.New(name)
		if err != nil {
			t.Fatalf("wasn't expecting an error, got %v", err)
		}

		if board.Name != name {
			t.Fatalf("expected %v, received %v", name, board.Name)
		}
	})

	t.Run("should return an error when name is empty", func(t *testing.T) {
		name := ""
		expected := board.ErrEmptyName
		_, actual := board.New(name)

		if expected != actual {
			t.Fatalf("expected %v, received %v", expected, actual)
		}
	})
}

func TestBoardAddTask(t *testing.T) {
	t.Run("should add a task to the board", func(t *testing.T) {
		b := newBoard(t)

		expected := 1
		tsk := newTask(t, 1)
		b.AddTask(tsk)

		if expected != len(b.Tasks) {
			t.Fatalf("expected %v, received %v", expected, len(b.Tasks))
		}
	})
}

func newBoard(t *testing.T) board.Board {
	t.Helper()

	board, err := board.New("board")
	if err != nil {
		t.Fatalf("wasn't expecting an error, got %v", err)
	}

	return board
}

func newTask(t *testing.T, n int) task.Task {
	t.Helper()

	title := fmt.Sprintf("new task - %v", n)
	task, err := task.New(title)
	if err != nil {
		t.Fatalf("wasn't expecting an error, got %v", err)
	}

	return task
}
