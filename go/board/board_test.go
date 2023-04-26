package board_test

import (
	"testing"

	"github.com/cesarFuhr/go/polyglot-todo/board"
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
