package structs

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type windowTemp struct {
	fun func(s string) string
}

var windowsT windowTemp

//var windowsParam = ""

type _utilites struct {
}

// проверяем есть ли впереди минус
func (u *_utilites) isMinus(str string) bool {
	if regexOne, err := regexp.Compile(`^-`); err == nil {
		return regexOne.MatchString(str)
	}
	return false
}

var utils _utilites

//var regexpMinus regexp.Regexp
var ErrUnknownСommands error

var printFormatRight []int
var printFormatLeft []int

// метод печати на экран
type printMode int

const (
	// не печатаем ничего, проверяем и отдаем ошибки
	PrintModeOnlyTest printMode = iota + 1
	// печатает Help
	PrintModeHelp
	// печатает ошибки если есть
	PrintModeShowError
	// печатает заполненые поля
	PrintModeShowValue
)
const (
	colorBlack   = "\033[1;30m"
	colorRed     = "\033[1;31m"
	colorGreen   = "\033[1;32m"
	colorYellow  = "\033[1;33m"
	colorPurple  = "\033[1;34m"
	colorMagenta = "\033[1;35m"
	colorTeal    = "\033[1;36m"
	colorWhite   = "\033[1;37m"
	colorEnd     = "\033[0m"
)

func init() {
	//utils := utilites{}

	// if regexOne, err := regexp.Compile(`^-`); err == nil {
	// 	regexpMinus = *regexOne
	// }
	ErrUnknownСommands = errors.New("неизвестные команды или флаги")

	printFormatRight = []int{15, 50}
	printFormatLeft = []int{-15, 50}
}

// var (
// 	Black   = Color("\033[1;30m%s\033[0m")
// 	Red     = Color("\033[1;31m%s\033[0m")
// 	Green   = Color("\033[1;32m%s\033[0m")
// 	Yellow  = Color("\033[1;33m%s\033[0m")
// 	Purple  = Color("\033[1;34m%s\033[0m")
// 	Magenta = Color("\033[1;35m%s\033[0m")
// 	Teal    = Color("\033[1;36m%s\033[0m")
// 	White   = Color("\033[1;37m%s\033[0m")
// )
func PrintConsole(color string, xx []int, str1 string, str string) {
	//tput cols  // ширина консоли
	if runtime.GOOS == "linux" {
		terminalCol := getWidthConsole() // ширина терминала
		fmt.Println(terminalCol)
	}

	var format string
	if color == "" {
		format = "% " + strconv.Itoa(xx[0]) + "s | % -" + strconv.Itoa(xx[1]) + "s\n"
	} else {
		format = "% " + strconv.Itoa(xx[0]) + "s " + color + "| % -" + strconv.Itoa(xx[1]) + "s" + colorEnd + "\n"
	}
	lenght := len(str)
	str = strings.ReplaceAll(str, "\n", " ")
	for {
		str = strings.ReplaceAll(str, "  ", " ")
		if lenght == len(str) {
			break
		}
		lenght = len(str)
	}

	words := strings.Split(str, " ")
	sumstr := ""
	x := xx[0]
	if x < 0 {
		x = x * -1
	}
	x = x + xx[1]
	//	x = xx[1]

	for _, word := range words {

		//	if (len(sumstr) + len(word)) > (xx[0] + xx[1]) {
		//len([]rune(russian)) кол-во русских букв без package utf8
		//utf8.RuneCountInString(sumstr)
		if (len([]rune(sumstr)) + len(word)) > x {
			fmt.Printf(format, str1, sumstr)
			str1 = ""
			sumstr = ""
		}
		sumstr += " " + word
		sumstr = strings.TrimSpace(sumstr)
	}
	if len(sumstr) > 0 {
		fmt.Printf(format, str1, sumstr)
	}
	//fmt.Printf("%s", "\n")
	// println(runtime.GOOS)
}

//var underline string

func (u *_utilites) setUnderline() string {
	x := printFormatLeft[0]
	if x < 0 {
		x = x * -1
	}
	arr := make([]string, x+1)
	return strings.Join(arr, "-")
}
