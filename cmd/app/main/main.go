// Main
package main

import (
	"fmt"
	_ "runtime"

	"github.com/ark-go/cli/internal"
	//"github.com/ark-go/cli/pkg/cli"
	//	_ "github.com/ark-go/cli/pkg/windows"
)

func main() {
	fmt.Println(".......... start .........\n ")
	cm := internal.CreateCmd()

	if err := cm.ParseCmdExitErrors(true); err != nil {
		//fmt.Printf(">>>>> %#v\n", err.(*cli.ErrorCli))
		println(">>", err.Error())
	}
	if m, err := cm.GetValues("addPath", "-p"); err == nil {
		for _, v := range m {
			println(v)
		}
	}
}
