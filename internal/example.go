package internal

import "github.com/ani-prog-go/cmdcli/pkg/structs"

func GetCommands() *structs.Commands {
	cmd := structs.Commands{
		structs.Command{
			Name:      "BBB",
			HelpShort: "Для теста",
			Flags: structs.Flags{

				structs.Flag{
					Name:      "lolololol",
					IsValue:   true,
					HelpShort: "этот ключ нужен для тестирования",
				},
				structs.Flag{
					Name:    "k",
					IsValue: true,
				},
				structs.Flag{
					Name:       "kmf",
					IsValue:    true,
					IsRequired: true,
				},
				structs.Flag{
					Name:    "notreb",
					IsValue: true,
				},
				structs.Flag{
					Name:    "notreb2",
					IsValue: true,
				},
			},
		},
		structs.Command{
			Name: "kkkk",
			Flags: structs.Flags{
				structs.Flag{
					Name:    "g",
					IsValue: true,
				},
				structs.Flag{
					Name:    "c",
					IsValue: true,
				},
			},
		},
	}
	return &cmd
}
