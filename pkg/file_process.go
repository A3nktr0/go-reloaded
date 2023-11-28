package go_reloaded

import (
	"bufio"
	"fmt"
	"os"
)

//************************************************************************************************
//************************************************************************************************
// Functions for file manipulation
//************************************************************************************************
//************************************************************************************************

// Read input file and return string
func ReadFile(input string) string {
	tmp_str := ""
	input_file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error during reading file")
	}
	defer input_file.Close()
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		tmp_str += string(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return tmp_str
}

// Write output file
func WriteFile(output, data string) {
	output_file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	output_file.WriteString(data)
	defer output_file.Close()
}
