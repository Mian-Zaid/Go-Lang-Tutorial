package Tasklistmanager

import (
	"bufio"
	"os"
)

type Task struct {
	id          int64
	description string
}

func (t *Task) inputTask() {
	scanner := bufio.NewScanner(os.Stdin)

	print("Enter new Task: ")
	if scanner.Scan() {
		t.description = scanner.Text()
	}
}
