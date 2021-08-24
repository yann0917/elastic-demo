package utils

import (
	"encoding/json"
	"fmt"
)

func PrintIndent(src interface{}) {
	fmt.Println("==========")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}