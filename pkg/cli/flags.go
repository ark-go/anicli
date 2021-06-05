package cli

//флаг
type flag struct {
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
func (f *flag) Required() *flag {
	f.isRequired = true
	return f
}

// Возвращает родительскую команду флага
func (f *flag) GetCommand() *command {
	return f.parent
}

// Устанавливает метку о том что значений у флага не ждем
func (f *flag) NoValues() *flag {
	f.noValues = true
	return f
}

// Добавляет флаг к команде
func (f *flag) AddFlag(name string, help string) *flag {
	return f.parent.AddFlag(name, help)
}

// Добавляет / изменяет справку о флаге
func (f *flag) AddHelp(help string) *flag {
	f.help = help
	return f
}

// Добавляет / заменяет короткую справку о флаге
func (c *flag) AddHelpShort(help string) *flag {
	c.helpShort = help
	return c
}

// устанавливаем метку о том что флаг был найден в коммандной строке
// func (f *flag) setPresent() *flag {
// 	f.isPresent = true
// 	return f
// }

// Возвращает значение и метку об установке значения
func (f *flag) GetValues() (bool, []string) {

	return f.isPresent, f.values
}
