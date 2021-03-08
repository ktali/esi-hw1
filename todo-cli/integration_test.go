// These tests require a running back-end service
package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/ktali/esi-hw1/todo-cli/cmd"
)

func TestAddNewTodo(t *testing.T) {
	content := []string{"testTodo"}
	cmd.AddTodo(content)
	if !strings.Contains(string(cmd.ListDos()), "testTodo") {
		t.Error("New todo was not found in the list of todos")
		return
	}
}

func TestRemoveTodo(t *testing.T) {
	todoContent := "remove test todo"
	content := []string{todoContent}
	cmd.AddTodo(content)
	dos := cmd.ListDos()
	todos := []cmd.Do{}
	json.Unmarshal(dos, &todos)

	var idToRemove []string

	for _, item := range todos {
		if item.Content == todoContent {
			idToRemove = []string{strconv.Itoa(int(item.Id))}
			break
		}
	}

	cmd.RemoveTodo(idToRemove)

	if strings.Contains(string(cmd.ListDos()), todoContent) {
		t.Error("Todo was not removed")
		return
	}
}

func TestCompleteTodo(t *testing.T) {
	todoContent := "Complete test todo"
	content := []string{todoContent}
	cmd.AddTodo(content)
	dos := cmd.ListDos()
	todos := []cmd.Do{}
	json.Unmarshal(dos, &todos)

	var idToComplete []string

	for _, item := range todos {
		if item.Content == todoContent {
			idToComplete = []string{strconv.Itoa(int(item.Id))}
			break
		}
	}

	cmd.CompleteTodo(idToComplete)

	newDos := cmd.ListDos()
	newTodos := []cmd.Do{}
	json.Unmarshal(newDos, &newTodos)

	var found bool = false

	for _, do := range newTodos {
		if do.Content == todoContent {
			found = true
			if !do.IsCompleted {
				t.Error("Todo was not completed")
			} 
			break;
		}
	}

	if !found {
		t.Error("Todo was not found")
	}
}

func TestRandomTodo(t *testing.T) {
	todoContent := "Random todo"
	content := []string{todoContent}
	cmd.AddTodo(content)
	do := cmd.GetRandom()

	if do == nil {
		t.Error("No TODO was returned")
	}
}

//Maybe not the best way of testing this. Oh well :)
func TestListTodos(t *testing.T) {
	//Get initial TODO list length
	dos := cmd.ListDos()
	todos := []cmd.Do{}
	json.Unmarshal(dos, &todos)

	initialLength := len(todos)

	//Add a new todo
	todoContent := "Test todo 1"
	content := []string{todoContent}
	cmd.AddTodo(content)

	//Check if the length is +1

	dos = cmd.ListDos()
	todos = []cmd.Do{}
	json.Unmarshal(dos, &todos)
	newLength := len(todos)

	if !(initialLength + 1 == newLength) {
		t.Error("New list length is not correct")
	}
}



