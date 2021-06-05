package cli

//флаг
type Flag struct {
	id         int      // для сортировки
	parent     *command // Пригодится для возврата родительской команды
	name       string
	isPresent  bool     // флаг найден/задан
	noValues   bool     // не будем проверять значения флага
	values     []string // если были данные
	isRequired bool     // флаг обязательный
	helpShort  string
	help       string
}

// Устанавливает флаг обязательный и возвращает родительскую команду
func (f *Flag) Required() *Flag {
	f.isRequired = true
	return f
}

// Возвращает родительскую команду флага
func (f *Flag) GetCommand() *command {
	return f.parent
}

// Устанавливает метку о том что значений у флага не ждем
func (f *Flag) NoValues() *Flag {
	f.noValues = true
	return f
}

// Добавляет флаг к команде
func (f *Flag) AddFlag(name string, help string) *Flag {
	return f.parent.AddFlag(name, help)
}

// Добавляет / изменяет справку о флаге
func (f *Flag) AddHelp(help string) *Flag {
	f.help = help
	return f
}

// Добавляет / заменяет короткую справку о флаге
func (c *Flag) AddHelpShort(help string) *Flag {
	c.helpShort = help
	return c
}

// устанавливаем метку о том что флаг был найден в коммандной строке
// func (f *Flag) setPresent() *Flag {
// 	f.isPresent = true
// 	return f
// }

// Возвращает значение и метку об установке значения
func (f *Flag) GetValues() (bool, []string) {

	return f.isPresent, f.values
}
