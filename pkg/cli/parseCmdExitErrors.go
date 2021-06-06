package cli

import (
	"os"
)

//	Разбирает командную строку,
//	 если чтото не сойдется вылетит с ошибкой и вернет ошибочные флаги
func (ac *AllCommands) ParseCmdExitErrors(printError bool) error {
	errorMsg := ErrorCli{}
	//musor := []string{}
	cmd := os.Args[1:]
	if len(cmd) == 0 {
		return nil
	}

	for i := 0; i < len(cmd); i++ {
		c := cmd[i]
		//var nocommand bool = false
		if i == 0 && utils.isMinus(c) && ac.Commands["Flags :"] != nil {
			// это первый элемент, это минус флаг, пустая команда есть в заявленных
			c = "Flags :"
			//	nocommand = true
		} else {
			if ac.Commands[c] == nil || ac.Commands[c].isPresent {
				// такой команды нет, или уже установлена  // а можно и не запрещать
				errorMsg.notCommand = true
				errorMsg.Command = c
				return &errorMsg
				// errUnknownСommands = errors.New("Данной команды [" + c + "] не существует")
				// return errUnknownСommands
			}

			// есть команда
			ac.Commands[c].isPresent = true

			if ac.Commands[c].noFlags || len(ac.Commands[c].flags) == 0 {
				// Флаги не нужны значит команда кончилась
				continue
			}

			//! Здесь начинается обработка флагов
			if (i + 1) >= len(cmd) {
				// кончилась командная строка, но мы сюда зачем то пришли, значит ошибка, ожидали флаг
				errorMsg.notFlag = true
				errorMsg.Command = c
				return &errorMsg
				// errUnknownСommands = errors.New("у данной команды [" + c + "] ожидался флаг")
				// return errUnknownСommands
			}

			i++ // преходим на ожидаемый флаг
			if !utils.isMinus(cmd[i]) {
				// оказалось это не флаг, ошибка
				errorMsg.notFlag = true
				errorMsg.Command = c
				return &errorMsg
				// errUnknownСommands = errors.New("у данной команды [" + c + "] ожидался флаг")
				// return errUnknownСommands
			}
		}
		// вероятно что есть флаг, пройдемся по строке в поиске всех флагов команды
		for (i) < len(cmd) {
			var flag string
			if !utils.isMinus(cmd[i]) { // НЕ начинается с "-"
				// это может сработать только после первого прохода, т.к. к for мы подходим с минусом
				// здесь без минуса вероятно команда, а должен быть флаг,
				i-- // вернем i на место и отправим проверять следующую, после флагов, команду
				break
			}

			flag = cmd[i] // получим название флага

			if ac.Commands[c].flags[flag] == nil {
				// У команды нет такого флага
				errorMsg.Command = c
				errorMsg.Flag = flag
				return &errorMsg
				// errUnknownСommands = errors.New("у данной команды [" + c + "] нет такого флага [" + flag + "]")
				// return errUnknownСommands
			}
			// Это наш флаг

			if ac.Commands[c].flags[flag].noValues {
				// флаг не требует значений, установим метку - есть в коммандной сроке
				i++ // фиксируем переход по строке
				ac.Commands[c].flags[flag].isPresent = true
				continue //продолжаем искать флаги
			}

			// до сюда дошли, значит, должно быть значение

			if (i + 1) >= len(cmd) {
				// закончились параметры, значения нет
				errorMsg.Command = c
				errorMsg.Flag = flag
				errorMsg.notVal = true
				return &errorMsg
				// errUnknownСommands = errors.New("Команда [" + c + "] у флага [" + flag + "] должно быть значение")
				// return errUnknownСommands
			}
			if utils.isMinus(cmd[i+1]) && !isNumber(cmd[i+1]) { // Начинается с "-" и не число
				// есть чтото с минусом но не цифра, а мы ждали значение
				errorMsg.Command = c
				errorMsg.Flag = flag
				errorMsg.notVal = true
				return &errorMsg
				// errUnknownСommands = errors.New("Команда [" + c + "] у флага [" + flag + "] должно быть значение")
				// return errUnknownСommands
				//break // ожидали значение без минуса, с минусом только если число, может команда начнем сначала
			}
			// значение есть, сохраним значение к флагу
			ac.Commands[c].flags[flag].isPresent = true // если бы у флага не было бы значения, вылетели бы раньше
			ac.Commands[c].flags[flag].values = append(ac.Commands[c].flags[flag].values, cmd[i+1])
			i++      // пропускаем флаг
			i++      // пропускаем значение
			continue // продолжаем искать флаги

		}
	}
	return ac.ParseCmdRequired()
}
