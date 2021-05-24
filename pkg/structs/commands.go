package structs

import (
	"errors"
	"fmt"
	"os"
)

// Все команды массив
type Commands []Command

type Command struct {
	//Название команды
	Name string
	//если команда была задана то true, поле скрыто
	isPresent bool
	//флаги команды []Flag
	Flags Flags
	//Требует обязательного наличия
	IsRequired bool
	//Справка для -help
	HelpShort string
	//Справка для -help
	Help     string
	Function func(f *Flags)
}

func (c *Commands) InitCommands() {
	for _, val := range *c {
		val.Flags.InitFlags()
	}

}

func (c *Commands) FindCommand(s string) (Command, error) {

	//for i, v := range *c {
	for i := 0; i < len(*c); i++ {
		if (*c)[i].Name == s {
			(*c)[i].isPresent = true
			return (*c)[i], nil
		}
	}
	return Command{}, errors.New("не найдено команды")
}

func (c *Commands) GetHelp() {
	if len(*c) > 0 {
		fmt.Printf("%s", "Требуемые и установленые параметры:\n")
	}
	for _, comm := range *c {
		if comm.isPresent { // соманда указана
			fmt.Printf("\033[1;32m%s\t\033[1;36m%s\033[0m\n", comm.Name, comm.HelpShort)
			for _, flag := range comm.Flags { // перебираем флаги

				if flag.isPresent || flag.IsRequired {
					//fmt.Printf("\t-%s - %s \n", flag.Name, flag.HelpShort)
					fmt.Printf("\033[1;32m\t-%s  \033[1;36m%s\033[0m\n", flag.Name, flag.HelpShort)
					if len(flag.Value) > 0 {
						fmt.Printf("\t  => %s", "")
						for _, valr := range flag.Value { // перебираем значения флага
							fmt.Printf(" '%s' ", valr)
						}
					} else {
						fmt.Printf("\033[1;31m\t  => %s\033[0m", "Этот флаг требует параметров!")
					}
					fmt.Printf("%s", "\n")
				}
				// else if !flag.isPresent {
				// 	fmt.Printf("\t  => %s\n", "false")
				// }
			}
		}
	}
}

/*
ParseCommand Ищет все команды (значения без минуса),
поле Flags каждой команды и оставшуюся часть коммандной строки args,
передает (ParseFlags) для поиска флагов, следующих за командой.

! Команда не найдет висящих флагов с минусом без впереди стоящей команды
*/
func (c *Commands) ParseCommands() {
	*NotFoundCommand = nil //new([]string)
	c.InitCommands()
	arrArg := os.Args
	// хранит адреса флагов и оставшуюся командную строку для разбора
	for i := 0; i < len(arrArg); i++ {
		if i == 0 {
			continue
		}
		val := arrArg[i]

		if matched := RegexpMinus.MatchString(val); !matched { // НЕ начинается с "-"
			if cm, err := c.FindCommand(val); err == nil {
				cm.Flags.ParseFlags(arrArg[i:])
			} else {
				*NotFoundCommand = append(*NotFoundCommand, val)
			}
		}
	}
	if len(*NotFoundCommand) > 0 {
		c.GetHelp()
		fmt.Printf("Ошибка: параметры не существуют, возможно у вас есть опечатки:\n%s\n", *NotFoundCommand)
		//	os.Exit(1)
	}
}
