package reader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/FlexK1ng/hw-basic/hw02_fix_app/types"
)

func ReadJSON(filePath string) (staff []types.Employee, err error) {
	var bytes []byte
	bytes, err = os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, nil
	}

	res := data

	return res, nil
}
