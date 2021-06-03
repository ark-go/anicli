package structs

import (
	"errors"
	"fmt"

	"strings"
)

func init() {

}
func (ac *AllCommands) PrintHelp(pm printMode) error {
	var checkedError string
	var errorParseCmd error
	var errorTest error
	if pm == PrintModeOnlyTest {
		if checkedError, errorParseCmd = ac.ParseCmd(true); errorParseCmd != nil {
			errorTest = ac.printCmd(pm, checkedError)
		}
	} else {
		ac.printCmd(pm, "")
	}
	return errorTest
}

func (ac *AllCommands) printCmd(pm printMode, checkedError string) error {
	if pm != PrintModeOnlyTest {
		if len(ac.HelpBefore) > 0 {
			fmt.Printf("%s\n", ac.HelpBefore)
		}
		fmt.Printf("%s\n%s\n", "Список команд:", utils.setUnderline())
	} else {
		if len(checkedError) > 0 {
			// if pm != PrintModeHelp {
			// 	PrintConsole("",printFormatRight, "мусор", checkedError+"\n")
			// }
		}
	}

	isReq := false
	notInstMess := ""
	requiredMess := ""
	keys := ac.sortedCommand()
	for _, v := range keys {
		// если команды нет, подготовим об этом сообщение
		if !ac.Commands[v].isPresent {
			notInstMess = "Not Present "
			if ac.Commands[v].isRequired {
				isReq = true
				requiredMess = "REQUIRED "
			}
		}
		if pm != PrintModeOnlyTest {
			PrintConsole("", printFormatLeft, ac.Commands[v].name, ac.Commands[v].helpShort)
			if pm == PrintModeShowValue {
				if len(requiredMess+notInstMess) > 0 {
					if len(requiredMess) > 0 {
						PrintConsole(colorRed, printFormatRight, "", "\b! "+requiredMess+notInstMess)
					} else {
						PrintConsole(colorPurple, printFormatRight, "", "\b> "+notInstMess)
					}
				}
			}
		}
		notInstMess = ""
		requiredMess = ""
		if len(ac.Commands[v].flags) > 0 { // если у команды есть ключи, прбежимся по ним
			keys2 := ac.Commands[v].sortedFlag() // сортируем ключи
			for _, v2 := range keys2 {
				// если флаг не установлен подготовим об этом сообщение
				if !ac.Commands[v].flags[v2].isPresent {
					notInstMess = "Not present "
					// флаг обязательный, только если команда установлена (обязательный для команды)
					if ac.Commands[v].flags[v2].isRequired && ac.Commands[v].isPresent {
						isReq = true
						requiredMess = "REQUIRED "
					}
				}
				// если не тест, то печатаем
				if pm != PrintModeOnlyTest {
					PrintConsole("", printFormatRight, ac.Commands[v].flags[v2].name, ac.Commands[v].flags[v2].helpShort)
					if pm == PrintModeShowValue { // показываем заданные значения
						if len(ac.Commands[v].flags[v2].values) > 0 {
							PrintConsole(colorTeal, printFormatRight, "", "\b> "+strings.Join(ac.Commands[v].flags[v2].values, ", "))
						} else if len(requiredMess) > 0 {
							PrintConsole(colorRed, printFormatRight, "", "\b> "+requiredMess+notInstMess)
						} else if len(notInstMess) > 0 {
							PrintConsole(colorPurple, printFormatRight, "", "\b> "+notInstMess)
						}

					}
				}
				notInstMess = ""
				requiredMess = ""
			}

		} else {
			// if pm != PrintModeOnlyTest && pm == PrintModeShowValue {
			// 	PrintConsole(colorPurple, printFormatRight, "", "Does not require flags") //"не требует значений")
			// }
		}
		if pm != PrintModeOnlyTest {
			fmt.Printf("\n") // пуревод строки между командами
		}
		//}

	}

	if isReq {
		{
			if pm != PrintModeOnlyTest && pm != PrintModeHelp {
				fmt.Printf("Не установлены обязательные значения\n")
			}
		}
		return errors.New("не установлены обязательные значения")
	}
	if pm != PrintModeOnlyTest {
		if len(ac.HelpAfter) > 0 {
			fmt.Printf("%s\n", ac.HelpAfter)
		}
	}
	return nil
}
