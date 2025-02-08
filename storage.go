package main

import (
	"encoding/json"
	"os"
)

func StoreLogs[T any](todo T) error {
	data, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		return err
	}

	if err = os.WriteFile("/home/kami-sama/projects/toDoListCli/json.txt", data, 0644); err != nil {
		return err
	}
	return nil
}

func LoadLogs[T any](todo *T) error {
	data, err := os.ReadFile("/home/kami-sama/projects/toDoListCli/json.txt")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, todo); err != nil {
		return err
	}
	return nil
}
