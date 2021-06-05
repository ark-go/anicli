// Все структуры проекта
package cli

import (
	"errors"
	"fmt"
	"sort"
)

type command struct {
	// автополе только для сортировки
	id int
	//Название команды
	name string
	//если команда была задана то true
	isPresent bool
	// не требует данных
	noFlags bool
	//Флаги этой команды
	flags map[string]*flag
	//Требует обязательного наличия, в командной строке
	isRequired bool
	//Короткая справка для -help
	helpShort string
	//Справка для -help
	help string
}

type AllCommands struct {
	Commands   map[string]*command
	HelpBefore string
	HelpAfter  string
	//	Rrrrr      type
}

// start
func GetCommands() *AllCommands {
	a := &AllCommands{}
	a.HelpBefore = ""
	a.HelpAfter = ""
	// a.printFormatRight = []int{15, 50}
	// a.printFormatLeft = []int{-15, 50}
	return a
}

/* Формат печати справки
   l - левое поле int
   r - правое поле int
*/
func (ac *AllCommands) SetPrintFormat(l int, r int) {
	printFormatRight = []int{l, r}     // левый отступ
	printFormatLeft = []int{l * -1, r} // размер правого поля
}

/* Проверка команды
если была в командной строке то вернет true
иначе error
*/
func (ac *AllCommands) IsCommand(command string) (bool, error) {
	com := ac.Commands[command]
	if com == nil {
		return false, errors.New("такой команды не существует, ошибка в имени команды")
	}
	if !com.isPresent {
		return false, errors.New("не установлена команда")
	}
	return true, nil

}

/* Добавить команду
   noFlags - флагов не будет
   name - имя команды
   help - краткая справка о команде
*/
func (ac *AllCommands) Add(name string, help string) *command {
	if ac.Commands == nil {
		ac.Commands = make(map[string]*command)
	}
	cmd := command{
		id:        len(ac.Commands),
		name:      name,
		noFlags:   false,
		helpShort: help,
		flags:     make(map[string]*flag),
	}
	if ac.Commands[name] != nil {
		fmt.Printf("WARNING: Повторно определена каманда: %s\n", name)
	}

	ac.Commands[name] = &cmd
	return &cmd
}

// Делает команду обязательной
func (c *command) Required() *command {
	c.isRequired = true
	return c
}

// У команды не будет флагов
func (c *command) NoFlags() *command {
	c.noFlags = true
	return c
}

// Добавляет / заменяет HelpShort команды
func (c *command) AddHelpShort(help string) *command {
	c.helpShort = help
	return c
}

// Добавляет / заменяет Help команды
func (c *command) AddHelp(help string) *command {
	c.help = help
	return c
}

/* Добавляет флаг к команде
   name - название флага
   help - краткая справка
*/
func (c *command) AddFlag(name string, help string) *flag {
	// if c.flags == nil {
	// 	c.flags = make(map[string]*flag)
	// }
	fl := &flag{
		id:        len(c.flags),
		noValues:  false,
		name:      name,
		helpShort: help,
		parent:    c,
	}
	if c.flags[name] != nil {
		fmt.Printf("WARNING: Повторно заявлен флаг: %s %s\n", c.name, name)
	}
	c.flags[name] = fl
	if c.noFlags {
		fmt.Printf("Запрашивается флаг, к команде не принимающих их %s %s\n", c.name, name)
	}
	return fl
}

// Возвращает значение и метку об установке значения
func (c *command) IsPresent() bool {
	return c.isPresent
}

// сортировка ключей Map command по значению id в command (автоматически назначаемому)
func (ac *AllCommands) sortedCommand() []string {

	keys := make([]string, 0, len(ac.Commands))
	for key := range ac.Commands {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if ac.Commands[keys[i]].id == ac.Commands[keys[j]].id {
			return ac.Commands[keys[i]].id < ac.Commands[keys[j]].id
		}
		return ac.Commands[keys[i]].id < ac.Commands[keys[j]].id
	})
	return keys
}

// сортировка ключей Map flags по значению id в Flag (автоматически назначаемому)
func (c *command) sortedFlag() []string {
	keys := make([]string, 0, len(c.flags))
	for key := range c.flags {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if c.flags[keys[i]].id == c.flags[keys[j]].id {
			return c.flags[keys[i]].id < c.flags[keys[j]].id
		}
		return c.flags[keys[i]].id < c.flags[keys[j]].id
	})
	return keys
}
