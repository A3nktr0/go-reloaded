package main

import (
	"fmt"
	fp "go_reloaded/pkg" // Import file_process's function
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) == 2 {
		data := fp.ReadFile(args[0])
		tmp_data := fp.SplitParenthesis(data)
		tmp_data = fp.Transform(tmp_data)
		tmp_data = fp.ReplaceA(tmp_data)
		tmp_data = fp.CheckComma(tmp_data)
		out_data := strings.Join(tmp_data, " ")
		out_data = fp.Punctuation(out_data)
		fp.WriteFile(args[1], out_data)
	} else {
		fmt.Println("Invalid input")
	}
}
