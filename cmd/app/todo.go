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

func (t *Todos) add(title string) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return
	}
	todo := Todo{
		Title:     title,
		Completed: false,
	}

	*t = append(*t, &todo)
}

// add another methode to Todos to check validity of index
func (t *Todos) CheckIndex(i int) error {

	if i < 0 || i >= len(*t) {
		return errors.New("Invalid index")
	}
	return nil
}
func (t *Todos) delete(tasks ...int) error {

	for _, i := range tasks {
		i--

		if err := t.CheckIndex(i); err != nil {
			return err
		}
		// *t = slices.Delete(*t, i, i+1)
		(*t)[i] = nil
	}
	cleanup(t)
	return nil
}
func (t *Todos) done(i int) error {
	i--
	if err := t.CheckIndex(i); err != nil {
		return err
	}

	(*t)[i].Completed = true

	return nil
}
func (t *Todos) log() {

	if len(*t) == 0 {
		fmt.Println("Your logs are empty..")
	}

	for i, strt := range *t {
		fmt.Printf("    %v > %v  ", i+1, strt.Title)
		if strt.Completed {
			fmt.Print("✔️  ☆*: .｡. o(≧▽≦)o .｡.:*☆")
		} else {
			fmt.Print("❌")
		}
		fmt.Print("\n")
	}
}
func (t *Todos) clear() {
	*t = Todos{}
}

func (t *Todos) purge() {
	for i := 0; i < len(*t); {
		if (*t)[i].Completed {
			*t = slices.Delete(*t, i, i+1)
		} else {
			i++
		}
	}
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
