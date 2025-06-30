package main

import (
	"flag"
	"fmt"
)

type cmdFlags struct {
	Add   string
	Del   int
	Done  int
	Log   bool
	Clear bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Adds a task")
	flag.IntVar(&cf.Del, "rm", -1, "Removes the task from log")
	flag.IntVar(&cf.Done, "done", -1, "Marks the task completed")
	flag.BoolVar(&cf.Log, "log", false, "Prints the task log")
	flag.BoolVar(&cf.Clear, "clear", false, "Clear all logs")

	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)
		args := flag.Args()
		for _, task := range args {
			todos.add(task)
		}
		if len(args) == 0 {
			fmt.Printf("   Task added to your log..\n\n")
		} else {
			fmt.Printf("   Tasks added to your log..\n\n")
		}
	case cf.Del != -1:
		todos.delete(cf.Del)
		fmt.Printf("   Task deleted from log..\n\n")
	case cf.Done != -1:
		todos.done(cf.Done)
		fmt.Printf("   Task completed..\n\n")
	case cf.Log:
		todos.log()
	case cf.Clear:
		todos.clear()
		fmt.Printf("    All logs cleared!\n\n")
	default:
		fmt.Printf("   $Invalid Command$\n\n")
	}
}
