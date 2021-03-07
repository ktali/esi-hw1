// These tests require a running back-end service
package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/ktali/esi-hw1/cli/cmd"
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
