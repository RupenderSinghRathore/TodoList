package main

func main() {
	todo := Todos{}
	LoadLogs(&todo)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todo)

	StoreLogs(todo)
}
