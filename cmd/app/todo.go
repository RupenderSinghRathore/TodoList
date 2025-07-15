package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Todo struct {
	Title     string
	Completed bool
}

type Todos []*Todo

func (todos *Todos) add(title string) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return
	}
	todo := Todo{
		Title:     title,
		Completed: false,
	}

	*todos = append(*todos, &todo)
}

// add another methode to Todos to check validity of index
func (todos *Todos) CheckIndex(i int) error {

	if i < 0 || i >= len(*todos) {
		return errors.New("Invalid index")
	}
	return nil
}
func (todos *Todos) delete(tasks ...int) error {

	for _, i := range tasks {
		i--

		if err := todos.CheckIndex(i); err != nil {
			return err
		}
		// *todos = slices.Delete(*todos, i, i+1)
		(*todos)[i] = nil
	}
	cleanup(todos)
	return nil
}
func (todos *Todos) done(i int) error {
	t := *todos
	i--
	if err := todos.CheckIndex(i); err != nil {
		return err
	}

	t[i].Completed = true

	return nil
}
func (todos *Todos) log() {

	if len(*todos) == 0 {
		fmt.Println("Your logs are empty..")
	}

	for i, strt := range *todos {
		fmt.Printf("    %v > %v  ", i+1, strt.Title)
		if strt.Completed {
			fmt.Print("✔️  ☆*: .｡. o(≧▽≦)o .｡.:*☆")
		} else {
			fmt.Print("❌")
		}
		fmt.Print("\n")
	}
}
func (todos *Todos) clear() {
	*todos = Todos{}
}

func cleanup(t *Todos) {
	for i := 0; i < len(*t); {
		if (*t)[i] == nil {
			*t = slices.Delete(*t, i, i+1)
			continue
		}
		i++
	}
}
