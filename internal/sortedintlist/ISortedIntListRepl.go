package sortedintlist

// ISortedIntListRepl - интерфейс командной консоли для работы с ISortedIntList
// в задаче было про то что есть логика, которая управляет этим списком
type ISortedIntListRepl interface {
	// PrintHelp Вывод справки
	PrintHelp()
	// Execute Полное выполнение консоли
	Execute()
	// ExecuteCommand Выполнение отдельной команды
	ExecuteCommand(command string) error
}