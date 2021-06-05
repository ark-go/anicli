// Main
package main

import (
	"fmt"
	_ "runtime"

	"github.com/ark-go/cli/internal"
	"github.com/ark-go/cli/pkg/cli"
	//	_ "github.com/ark-go/cli/pkg/windows"
)

func main() {
	fmt.Println(".......... start .........\n ")
	cm := internal.CreateCmd()

	// if _, err := cm.ParseCmd(true); err != nil {
	// 	if err == structs.ErrUnknownСommands {
	// 		cm.PrintHelp(structs.PrintModeOnlyTest) // OnlyTest печатает еррор
	// 	}
	// }
	// if err := cm.PrintHelp(structs.PrintModeOnlyTest); err != nil {
	// 	cm.PrintHelp(structs.PrintModeShowValue)
	// }
	if err := cm.ParseCmdExitErrors(true); err != nil {
		//if err == structs.ErrUnknownСommands {

		println(">>>>>" + err.(*structs.ErrorCli).Command)
		println(">>")
		//println(">>", err(m.Key()))
		// if r, ok := err.(cm.ErrorCli); ok {

		// }
		//	println("KK>", cm.ErrorCli.Error())
		//}
	}
	// if err := cm.ParseCmdRequired(); err != nil {
	// 	println(">>", err.Error())
	// }
}
