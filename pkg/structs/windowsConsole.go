// +build windows

// сработает только при сборке виндовс
// GOOS=windows GOARCH=amd64 go build -o ./bin/main/winmain.exe cmd/app/main/main.go
package structs

import (
	"golang.org/x/sys/windows"
	"os"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32
	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

}
