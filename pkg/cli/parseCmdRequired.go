package cli

type ErrorCli struct {
	TestRequired bool
	Command      string
	Flag         string
	notCommand   bool
	notFlag      bool
	notVal       bool
}

func (e *ErrorCli) Error() string {
	if e.TestRequired {
		if e.notCommand {
			mess := "*Нeт обязательной команды " + e.Command
			e.notCommand = false
			e.Command = ""
			return mess
		}
		if e.notFlag {
			mess := "*нет флага " + e.Flag + " у команды " + e.Command
			e.notFlag = false
			e.Command = ""
			e.Flag = ""
			return mess
		}
		return "Ошибка раз"
	} else {
		if e.notFlag {
			mess := "The " + e.Command + " command requires flags"
			e.notFlag = false
			e.Command = ""
			//! Для команды требуются флаги
			return mess
		}
		if e.notVal {
			mess := "For command " + e.Command + ", the " + e.Flag + " flag must have a value"
			e.notVal = false
			e.Command = ""
			e.Flag = ""
			//! Для команды XXX, флаг YYY должен иметь значение
			return mess
		}
		if e.Command != "" && e.Flag != "" {
			mess := "For command " + e.Command + ", the " + e.Flag + " flag doesn't exist"
			e.Command = ""
			e.Flag = ""
			//! Для команды XXX флаг YYY не существует
			return mess
		}
		if e.Command != "" {
			mess := "The " + e.Command + " command doesn't exist"
			e.Command = ""
			//! команда не существует
			return mess
		}
	}
	return "error cli"
}

func (ac *AllCommands) ParseCmdRequired() error {
	keys := ac.sortedCommand()
	for _, key := range keys {

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
			// если требуется и установлена, проверим флаги
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
			} else {
				//! это все лишнее
				// команда присутствует
				if !val.noValues && len(val.values) == 0 {
					// Значение должно быть, но его нет
					return &ErrorCli{
						TestRequired: true,
						Command:      c.name,
						Flag:         val.name,
						notVal:       true,
					}
				}
			}
		}
	}
	return nil

}
