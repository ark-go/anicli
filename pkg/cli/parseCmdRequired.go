package cli

// Експортируется клиенту, малоли ему надо, для приведения error ответа
//  к структуре ErrorCli
type ErrorCli struct {
	TestRequired bool   // тест на обязательные true / установленные false команды-флаги
	Command      string // команда
	Flag         string // флаг
	notCommand   bool   // нет команды
	notFlag      bool   // нет флага
	notVal       bool   // нет значения
}

func (e *ErrorCli) Error() string {
	if e.TestRequired {
		if e.notCommand {
			//!Нeт обязательной команды
			return "*No mandatory " + e.Command + " command"
		}
		if e.notFlag {
			//! нет обязательного флага у команды
			if e.Command != "Flags :" {
				return "*No mandatory " + e.Flag + " flag, commands " + e.Command
			} else {
				return "*No mandatory " + e.Flag + " flag"
			}
		}
		return "Unknown error while checking required flags"
	} else {
		if e.notFlag {
			//! Для команды требуются флаги
			return "The " + e.Command + " command requires flags"
		}
		if e.notVal {
			//! Для команды XXX, флаг YYY должен иметь значение
			if e.Command != "Flags :" {
				return "For command " + e.Command + ", the " + e.Flag + " flag must have a value"
			} else {
				return "The " + e.Flag + " flag must have a value"
			}
		}
		if e.Command != "" && e.Flag != "" {
			//! Для команды XXX флаг YYY не существует
			if e.Command != "Flags :" {
				return "For command " + e.Command + ", the " + e.Flag + " flag doesn't exist"
			} else {
				return "The " + e.Flag + " flag doesn't exist"
			}
		}
		if e.Command != "" {
			//! команда не существует
			return "The " + e.Command + " command doesn't exist"
		}
	}
	return "Unknown error while checking command line"
}

// тест на несуществующие ключи
// если они существуют, они проидут тест раньше в parseCmdExitErrors
func (ac *AllCommands) ParseCmdRequired() error {
	keys := ac.sortedCommand()
	for _, key := range keys {
		// if key == "Flags :" {
		// 	continue
		// }
		if ac.Commands[key].isRequired {
			// если команда требуется
			if !ac.Commands[key].isPresent {
				// если команды нет
				return &ErrorCli{
					TestRequired: true,
					Command:      key,
					notCommand:   true,
				}

			}
		}
		if ac.Commands[key].isPresent {
			// если установлена, проверим флаги
			if err := ac.parseFlagsRequired(ac.Commands[key]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (co *AllCommands) parseFlagsRequired(c *command) error {

	for _, val := range c.flags {

		if val.isRequired {
			// флаг обязательный
			if !val.isPresent {
				// флаг не указан
				return &ErrorCli{
					TestRequired: true,
					Command:      c.name,
					Flag:         val.name,
					notFlag:      true,
				}
			}
		}
	}
	return nil

}
