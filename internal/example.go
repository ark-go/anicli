// Пакет с примером для ...
package internal

import "github.com/ark-go/cli/pkg/structs"

func CreateCmd() *structs.AllCommands {
	cmd := structs.GetCommands()
	cmd.HelpBefore = `
Пример пакета, позволяющего разбирать коммандную строку и создавать заранее известную структуру команд и флагов,
правила простые...
`
	cmd.HelpAfter = `copyright 2021
`
	cmd.Add("help", "Вывод справки").NoFlags().Required()
	cmd.Add("-help", "Вывод справки")
	cmd.Add("copy", "Копирование файла").
		AddFlag("-r", "путь к файлу для чтения, если путь с пробелами то в кавычках").Required().
		AddFlag("-w", "Путь к файлу для записи, если путь с пробелами то в кавычках")
	cmd.Add("reverse", "копирует и перефорачивает файл задом наперед")
	cmd.Add("addPath", "Добавить дополнительные пути для копирования файла").
		AddFlag("-p", "Путь к дополнительному файлу, флаг может повторятся для нескольких путей")
	return cmd
}
