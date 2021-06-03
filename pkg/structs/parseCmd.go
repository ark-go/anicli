package structs

import (
	"os"
	"strconv"
	"strings"
)

func isNumber(str string) bool {
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		return true
	}
	return false
}

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
		if ac.Commands[c] == nil {
			// такой команды нет
			musor = append(musor, c)
			continue
		}
		// есть команда
		ac.Commands[c].isPresent = true
		// Это bool значит команда кончилась
		if ac.Commands[c].noFlags {
			continue
		}
		// Есть следующий параметр
		if (i + 1) < len(cmd) {
			for (i + 1) < len(cmd) {
				var flag string
				//	if matched := regexpMinus.MatchString(cmd[i+1]); !matched { // НЕ начинается с "-"
				if !utils.isMinus(cmd[i+1]) { // НЕ начинается с "-"
					// здесь без минуса вероятно команда, а должен быть флаг, запускаем обратно в цикл команд
					break
				} else {
					// тут минус впереди, идут флаги
					// если нет такого флага у нас то идем по cmd дальше
					flag = cmd[i+1] // получим название флага
					if ac.Commands[c].flags[flag] == nil {
						// TODO: флаг левый  обработать
						musor = append(musor, flag) // собираем левые флаги (у данной команды нет такого флага)
						i++
						continue // нет такоо флага идем по cmd дальше

					}
					// флаг без значений?
					if ac.Commands[c].flags[flag].noValues {
						i++
						// флаг найден, он не требует значений, установим метку - есть в коммандной сроке
						ac.Commands[c].flags[flag].isPresent = true
						continue //продолжаем флаги
					}

					//ac.Commands[c].flags[flag].isPresent = true // сам флаг есть но..
					i++
					// если впереди в CMD еще чтото есть
					if (i + 1) < len(cmd) {

						//	if matched := regexpMinus.MatchString(cmd[i+1]); matched && !isNumber(cmd[i+1]) { // Начинается с "-"
						if utils.isMinus(cmd[i+1]) && !isNumber(cmd[i+1]) { // Начинается с "-" и не число
							//ac.Commands[c].flags[flag].isPresent = false
							break // ожидали значение без минуса, может команда начнем сначала
						} else {
							// сохраним значение к флагу
							ac.Commands[c].flags[flag].values = append(ac.Commands[c].flags[flag].values, cmd[i+1])
							ac.Commands[c].flags[flag].isPresent = true
							i++ //i + 2
							continue
						}
					}

					continue // крутим i3 флаги
				}
			}
		} else {
			break // кончились команды выходим из For
		}
	}
	if len(musor) > 0 {
		// if printError {
		// 	fmt.Printf("Неизвестные команды или флаги:\n")
		// 	PrintConsole("",printFormatRight, "мусор", strings.Join(musor, ", ")+"\n")
		// }
		return strings.Join(musor, ", "), ErrUnknownСommands
	}

	return "", nil
}
