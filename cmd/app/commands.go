package main

import (
	"flag"
	"fmt"
	"strconv"
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
		todos.log()
	case cf.Del != -1:
		args := flag.Args()
		lis := []int{cf.Del}
		for _, s := range args {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("   > %s has to be a positive Integer\n", s)
			}
			lis = append(lis, i)
		}
		todos.delete(lis...)
		if len(args) == 0 {
			fmt.Printf("   Task deleted from log..\n\n")
		} else {
			fmt.Printf("   Tasks deleted from log..\n\n")
		}
		todos.log()
	case cf.Done != -1:
		todos.done(cf.Done)
		args := flag.Args()
		for _, taskstr := range args {
			task, err := strconv.Atoi(taskstr)
			if err != nil {
				fmt.Printf("   > %s has to be a positive Integer\n", taskstr)
			}
			todos.done(task)
		}
		if len(args) == 0 {
			fmt.Printf("   Task completed..\n\n")
		} else {
			fmt.Printf("   Tasks completed..\n\n")
		}
		todos.log()
	case cf.Log:
		todos.log()
	case cf.Clear:
		todos.clear()
		fmt.Printf("    All logs cleared!\n\n")
	default:
		fmt.Printf("   $Invalid Command$\n\n")
	}
}
