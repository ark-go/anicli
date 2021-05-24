package structs

import (
	"errors"
	"regexp"
)

//флаги массив
type Flags []Flag
type Flag struct {
	Name string
	// флаг найден
	isPresent bool
	IsValue   bool //  у флага ждать значение
	Value     []string
	ValType   interface{}
	// флаг обязательный
	IsRequired bool
	HelpShort  string
	Help       string
}

func (s *Flags) InitFlags() {
	for i := 0; i < len(*s); i++ {
		(*s)[i].Value = []string{}
	}
}
func (f *Flags) FindFlags(s string) (*Flag, error) {

	if matched, err := regexp.MatchString(`^-{1}\w`, s); err == nil && matched {
		newS := s[1:] // убираем минус
		for i := 0; i < len(*f); i++ {
			//	for _, v := range *f {
			if (*f)[i].Name == newS {
				(*f)[i].isPresent = true

				return &(*f)[i], nil
			}
		}

	}

	return nil, errors.New("не найдено флага")
}

// flg = s
func (s *Flags) ParseFlags(args []string) {
	var bufferFlag *Flag // храним адрес флага для записи в него значения если оно есть

	for i, val := range args {
		if i < 1 {
			continue
		}
		var cm *Flag
		if matched := RegexpMinus.MatchString(val); matched { // начинается с минуса "-"
			bufferFlag = nil // получили флаг,значит очистим буфер для адреса флага
			if cm, _ = s.FindFlags(val); cm != nil {
				if cm.IsValue { //! Эта проверка вроде лишняя, параметр говорит о том, что мы ожидаем параметр
					// Ожидаем что у флага будет значение, сохраняем адрес флага
					bufferFlag = cm // на следущем шаге, если в начале нет минуса, заберем значение
				}
			} else {
				*NotFoundFlag = append(*NotFoundFlag, val) // флага с таким именем (val) нет
			}
		} else {
			if bufferFlag != nil { // ожидали значение
				bufferFlag.Value = append(bufferFlag.Value, val) // отправляем значение по адресу
			} else {
				break // кончились флаги у команды, т.е. началась новая команда - выходим из for
			}

		}
	}
}
