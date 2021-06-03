// Main
package main

import (
	"fmt"
	_ "runtime"

	"github.com/ark-go/cli/internal"
	"github.com/ark-go/cli/pkg/structs"
	//	_ "github.com/ark-go/cli/pkg/windows"
)

func main() {
	// if runtime.GOOS == "windows" {
	// 	ark.InitWindows()
	// }
	// у команды свои флаги, они считываются до следущей команды
	// команды могут быть с ведущим минусом !?
	fmt.Println(".......... start .........\n ")
	//all := structs.AllCommands{}
	// all := structs.GetCommands()

	// //all.SetPrintFormat(30, 50)
	// tst := all.Add("-go", "Команда для старта").Required().AddFlag("-qwerty", "Это м флаг").
	// 	AddFlag("-vvv", "Описание флага").Required().
	// 	AddFlag("-vvv", "Описание флага22")
	// tst.AddFlag("-qwerty2", "Второй").NoValues()
	// all.Add("-go", "Команда для старта").Required().AddFlag("-qwerty789", "Это м флаг")
	// all.Add("-BBB", "Команда для старта").Required().AddFlag("-lolololol", "Это м флаг")
	// all.Add("MMM", "Команда для старта").AddFlag("-ku", "Это м флаг")
	// all.Add("BBB", "Команда для старта").Required().AddFlag("-lolololol", "Это м флаг")

	// cmd := all.Add("TEST", "тест команда").AddFlag("-c", "Пример с")
	// all.Add("TEST888", "тест команда").AddFlag("-c", "Пример с")
	// cmd.AddFlag("-g", "Укажите этот флаг")
	// cmd.AddFlag("-sss", "Укажите этот Bool флаг")

	// распарсиваем командную строку по нашим структурам

	cm := internal.CreateCmd()

	// if _, err := cm.ParseCmd(true); err != nil {
	// 	if err == structs.ErrUnknownСommands {
	// 		cm.PrintHelp(structs.PrintModeOnlyTest) // OnlyTest печатает еррор
	// 	}
	// }
	if err := cm.PrintHelp(structs.PrintModeOnlyTest); err != nil {
		cm.PrintHelp(structs.PrintModeShowValue)
	}
}
