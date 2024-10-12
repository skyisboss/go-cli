package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Println(a ...any) {
	fmt.Println("🟥🟩🟨")
	fmt.Println("")
	fmt.Println(a...)
	fmt.Println("")
	fmt.Println("🟥🟩🟨")
}

// 美化打印
func ToJson(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}
