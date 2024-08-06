package init

import (
	"fmt"

	"github.com/FlexK1ng/hw-basic/hw02_fix_app/printer"
	"github.com/FlexK1ng/hw-basic/hw02_fix_app/reader"
	"github.com/FlexK1ng/hw-basic/hw02_fix_app/types"
)

func init() {
	var path string = "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
