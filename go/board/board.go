package board

import (
	"time"

	"github.com/cesarFuhr/go/polyglot-todo/task"
)

type Board struct {
	Name      string
	Tasks     []task.Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
