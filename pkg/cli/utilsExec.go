//

package cli

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getWidthConsole() int {
	cmd := exec.Command("tput", "cols")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		//	fmt.Printf(err.Error())

		if l, err := strconv.Atoi(strings.TrimSpace(out.String())); err == nil {
			return l
		} else {
			fmt.Printf("%s", err.Error())
		}
		//fmt.Printf("in all caps: %q\n", out.String())
	}

	//cmd := exec.Command("tput", "cols")

	// Stdout + stderr
	// out, err := cmd.StderrPipe() // rm writes the prompt to err
	// if err != nil {
	// 	println(err)
	// 	return 0
	// }
	// r := bufio.NewReader(out)

	// // Stdin
	// in, err := cmd.StdinPipe()
	// if err != nil {
	// 	println(err)
	// 	return 0
	// }
	// defer in.Close()

	// // Start the command!
	// err = cmd.Start() //.Start()
	// if err != nil {
	// 	println(err)
	// 	return 0
	// }

	// if line, _, err := r.ReadLine(); err == nil {

	// 	fmt.Printf("%s", string(line))

	// 	if l, err := strconv.Atoi(string(line)); err != nil {
	// 		return l
	// 	}
	// } else {
	// 	fmt.Printf(">>%s", err.Error())
	// }
	return 0
	// for err != nil {
	// 	if string(line) == "Remove file 'somefile.txt'?" {
	// 		in.Write([]byte("y\n"))
	// 	}
	// 	line, _, err = r.ReadLine()
	// }
}
