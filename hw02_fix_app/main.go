package main

import (
	"fmt"
	"os"

	"github.com/FlexK1ng/hw-basic/hw02_fix_app/printer"
	"github.com/FlexK1ng/hw-basic/hw02_fix_app/reader"
	"github.com/FlexK1ng/hw-basic/hw02_fix_app/types"
)

func main() {
	defaultpath := "data.json"
	var path string
	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = defaultpath
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	printer.PrintStaff(staff)
}
