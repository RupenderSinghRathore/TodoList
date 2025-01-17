package main

import (
	"errors"
	"fmt"
)

type Todo struct{
    Title string
    Completed bool
}

type Todos []Todo 

func (todos *Todos) add(title string) {
    todo := Todo{
            Title: title,
            Completed: false,
        }

    *todos  = append(*todos, todo) 
}
// add another methode to Todos to check validity of index
func (todos *Todos) CheckIndex(i int) error {

    if i < 0 || i >= len(*todos) {
        return errors.New("Invalid index")
    }
    return nil
}
func (todos *Todos) delete(i int) error {

    i--

    if err := todos.CheckIndex(i); err != nil {
        return err
    }
    t := *todos
    *todos = append(t[:i], t[i+1:]...)
    return nil
}
func (todos *Todos) done(i int) error {
    t := *todos
    i--
    if err := todos.CheckIndex(i); err != nil {
        return err
    }

    strt := &t[i]
    strt.Completed = true

    return nil
}
func (todos *Todos) log() {
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





    

