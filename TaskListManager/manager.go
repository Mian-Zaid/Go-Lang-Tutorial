package Tasklistmanager

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Manager struct {
	File     *os.File
	filePath string
}

func (m *Manager) AddTask() {
	task := Task{}
	task.inputTask()

	m.WriteToFile(m.File, task.description)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// create file
func (m *Manager) CreateOrOpenTextFile(filename string) *os.File {
	if fileExists(filename) {
		return m.OpenFile(filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil
	}
	return file

}

// open file
func (m *Manager) OpenFile(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	return file
}

// close file
func CloseFile(file *os.File) {
	file.Close()
}

// write to file
func (m *Manager) WriteToFile(file *os.File, data string) {
	if file != nil {
		_, err := file.WriteString(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

// read from file
func ReadFile(file *os.File) []byte {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return content
}

//edit task

//delete task

//view task

//save task to file

//read task from file
