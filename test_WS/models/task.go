package models

import (
	"errors"
	"fmt"
)

type Task struct {
	ID          int
	Description string
	Progress    float32
	IsDeleted   bool
}

var (
	tasks      []*Task
	nextTaskID = 1
)

func GetTasks() []*Task {
	var aux []*Task
	for _, t := range tasks {
		if !t.IsDeleted {
			aux = append(aux, t)
		}
	}
	return aux
}

func AddTask(t Task) (Task, error) {
	if t.ID != 0 {
		return Task{}, errors.New("New Task must not include id or it must be set to zero")
	}
	if t.Progress > 1 || t.Progress < 0 {
		return Task{}, errors.New("New Task Progress must be between 0 and 1")
	}
	t.ID = nextTaskID
	nextTaskID++
	tasks = append(tasks, &t)
	return t, nil
}

func GetTaskByID(id int) (Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return *t, nil
		}
	}

	return Task{}, fmt.Errorf("Task with ID '%v' not found", id)
}

func UpdateTask(t Task) (Task, error) {
	for i, candidate := range tasks {
		if candidate.ID == t.ID {
			tasks[i] = &t
			return t, nil
		}
	}

	return Task{}, fmt.Errorf("Task with ID '%v' not found", t.ID)
}

func RemoveTaskById(id int) error {
	for _, t := range tasks {
		if t.ID == id {
			t.IsDeleted = true
			return nil
		}
	}

	return fmt.Errorf("Task with ID '%v' not found", id)
}
