package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveTaskById(t *testing.T) {

	var mockTask *Task = new(Task)

	mockTask.ID = 1
	mockTask.Description = "desc"
	mockTask.Progress = 0.1
	mockTask.IsDeleted = false

	tasks = append(tasks, mockTask)

	id := 1
	if RemoveTaskById(id) != nil {
		t.Error("Do not returned nil")
	} else {
		if !tasks[0].IsDeleted {
			t.Error("isDeleted should be true")
		} else {
			t.Logf("test1 passed")
		}
	}

	result := RemoveTaskById(2)

	expectedMessage := fmt.Errorf("Task with ID '%v' not found", 2)
	if !reflect.DeepEqual(result, expectedMessage) {
		t.Error("message not as expected, expected: ", expectedMessage)
	} else {
		t.Logf("test2 passed")
	}

}
