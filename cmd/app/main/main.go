package main

import (
	"fmt"

	"github.com/ark-go/anicli/internal"
)

func main() {
	// 1 - если начинается со флагов, то они считываются до первой команды
	// 2 - у команды свои флаги, они считываются до следущей команды
	fmt.Println(".......... start .........")
	// Получаем, зачем ??
	cmd := internal.GetCommands()
	// Проверка и сопоставление флагов
	cmd.ParseCommands()
	cmd.ParseCommands()
	cmd.ParseCommands()
	// Справка
	//cmd.GetHelp()
}
