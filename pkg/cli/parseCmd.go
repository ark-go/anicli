package cli

import (
	"os"
	"strings"
)

//  Разбирает командную строку, при ошибках вернет ошибочные флаги
// 		printError - распечатет на консоли
func (ac *AllCommands) ParseCmd(printError bool) (string, error) {
	musor := []string{}
	cmd := os.Args[1:]
	if len(cmd) == 0 {
		return "", nil
	}

	for i := 0; i < len(cmd); i++ {
		c := cmd[i]
		if ac.Commands[c] == nil || ac.Commands[c].isPresent {
			// такой команды нет, или уже установлена  // а можно и не запрещать
			musor = append(musor, c)
			continue
		}
		// есть команда
		ac.Commands[c].isPresent = true

		if ac.Commands[c].noFlags {
			// Это флаги не нужны значит команда кончилась
			continue
		}

		if (i + 1) >= len(cmd) {
			// кончилась командная строка выходим из For
			break
		}
		// вероятно что есть флаги
		for (i + 1) < len(cmd) {
			var flag string
			if !utils.isMinus(cmd[i+1]) { // НЕ начинается с "-"
				// здесь без минуса вероятно команда, а должен быть флаг, запускаем обратно в цикл команд
				break
			} else {
				// этофлаг с минусом
				flag = cmd[i+1] // получим название флага
				if ac.Commands[c].flags[flag] == nil {
					// TODO: флаг левый  обработать
					musor = append(musor, flag) // собираем левые флаги (у данной команды нет такого флага)
					i++                         // фиксируем переход по строке
					continue                    // нет такого флага идем по cmd дальше

				}
				// Это наш флаг?
				if ac.Commands[c].flags[flag].noValues {
					// флаг не требует значений, установим метку - есть в коммандной сроке
					i++ // фиксируем переход по строке
					ac.Commands[c].flags[flag].isPresent = true
					continue //продолжаем искать флаги
				}
				// должно быть значение
				//ac.Commands[c].flags[flag].isPresent = true //???  флаг вроде есть, мы не знаем есть ли значения
				i++ // фиксируем переход по строке, переходим за флаг
				// если впереди в CMD еще чтото есть
				if (i + 1) >= len(cmd) {
					break // командная строка закончилась, а значения флага нет
				}

				//	if matched := regexpMinus.MatchString(cmd[i+1]); matched && !isNumber(cmd[i+1]) { // Начинается с "-"
				if utils.isMinus(cmd[i+1]) && !isNumber(cmd[i+1]) { // Начинается с "-" и не число
					break // ожидали значение без минуса, с минусом только если число, может команда начнем сначала
				}
				// значение есть, сохраним значение к флагу
				ac.Commands[c].flags[flag].values = append(ac.Commands[c].flags[flag].values, cmd[i+1])
				ac.Commands[c].flags[flag].isPresent = true //???
				i++                                         //i + 2 // фиксируем переход по строке, переходим за значение
				continue                                    // продолжаем искать флаги
			}
		}
	}
	if len(musor) > 0 {
		// if printError {
		// 	fmt.Printf("Неизвестные команды или флаги:\n")
		// 	PrintConsole("",printFormatRight, "мусор", strings.Join(musor, ", ")+"\n")
		// }
		return strings.Join(musor, ", "), errUnknownСommands
	}

	return "", nil
}
