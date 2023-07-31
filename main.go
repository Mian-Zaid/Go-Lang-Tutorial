package main

import (
	tasklistmanager "gotutorial.com/main/TaskListManager"
)

func main() {
	manager := &tasklistmanager.Manager{}
	const data = "Hello this is text"

	//create DB file if not exists
	const filePath = "TaskListManager/Database/db.txt"
	manager.File = manager.CreateOrOpenTextFile(filePath)

	manager.AddTask()

}
