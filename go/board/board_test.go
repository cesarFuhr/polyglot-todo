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

func TestBoardInsertTask(t *testing.T) {
	t.Run("should add a task to the board", func(t *testing.T) {
		b := newBoard(t)

		expected := 1
		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}
	})

	t.Run("should add a task to the top of the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		tsk = newTask(t, 2)
		b.InsertTask(0, tsk)

		expected := 2

		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		actualTask, _ := b.GetTask(0)
		if actualTask.Title != tsk.Title {
			t.Fatalf("expected title %v, received title %v", b.GetTasks()[0].Title, tsk.Title)
		}
	})

	t.Run("should add a task to the end of the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		tsk = newTask(t, 2)
		b.InsertTask(1, tsk)

		expected := 2

		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		expectedTitle := tsk.Title
		actualTask, _ := b.GetTask(1)
		if actualTask.Title != expectedTitle {
			t.Fatalf("expected title %v, received title %v", expectedTitle, actualTask.Title)
		}
	})

	t.Run("should add a task to the middle of the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		tsk = newTask(t, 2)
		b.InsertTask(1, tsk)

		middleTask := newTask(t, 3)
		b.InsertTask(1, middleTask)

		expected := 3

		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		expectedTitle := middleTask.Title
		actualTask, _ := b.GetTask(1)
		if actualTask.Title != expectedTitle {
			t.Fatalf("expected title %v, received title %v", expectedTitle, b.GetTasks()[1].Title)
		}
	})
}

func TestBoardRemoveTask(t *testing.T) {
	t.Run("should remove single task from the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		b.RemoveTask(0)

		expected := 0
		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}
	})

	t.Run("should remove task at the end for the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		tsk2 := newTask(t, 2)
		b.InsertTask(1, tsk2)

		b.RemoveTask(1)

		expected := 1
		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		expectedTitle := tsk.Title
		actualTask, _ := b.GetTask(0)
		if actualTask.Title != expectedTitle {
			t.Fatalf("expected title %v, received title %v", expectedTitle, actualTask.Title)
		}
	})

	t.Run("should remove task at the beginning for the board", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		tsk2 := newTask(t, 2)
		b.InsertTask(1, tsk2)

		b.RemoveTask(0)

		expected := 1
		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		expectedTitle := tsk2.Title
		actualTask, _ := b.GetTask(0)
		if actualTask.Title != expectedTitle {
			t.Fatalf("expected title %v, received title %v", expectedTitle, actualTask.Title)
		}
	})
}

func TestBoardGetTask(t *testing.T) {
	t.Run("should get a task", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		b.GetTask(0)

		expected := 1
		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		actualTask, _ := b.GetTask(0)
		if actualTask.Title != tsk.Title {
			t.Fatalf("expected title %v, received title %v", tsk.Title, actualTask.Title)
		}
	})

	t.Run("should return an invalid position error", func(t *testing.T) {
		b := newBoard(t)

		_, err := b.GetTask(1)

		expected := board.ErrInvalidPosition

		if expected != err {
			t.Fatalf("expected %v, received %v", expected, err)
		}
	})

}

func TestBoardUpdateTask(t *testing.T) {
	t.Run("should update task", func(t *testing.T) {
		b := newBoard(t)

		tsk := newTask(t, 1)
		b.InsertTask(0, tsk)

		expectedTitle := "title updated"
		b.UpdateTask(0, task.Task{Title: expectedTitle})

		expected := 1
		if expected != len(b.GetTasks()) {
			t.Fatalf("expected %v, received %v", expected, len(b.GetTasks()))
		}

		actualTask, _ := b.GetTask(0)
		if actualTask.Title != expectedTitle {
			t.Fatalf("expected title %v, received title %v", expectedTitle, actualTask.Title)
		}
	})

	t.Run("should return an invalid position error", func(t *testing.T) {
		b := newBoard(t)

		err := b.UpdateTask(1, task.Task{})

		expected := board.ErrInvalidPosition

		if expected != err {
			t.Fatalf("expected %v, received %v", expected, err)
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
