package main

import (
	"fmt"
	"os"
	"strings"

	"wogra.com/executor"
)

func runAsLocalExecutor(args []string) bool {
	if len(args) >= 2 && args[0] == "-execute" {
		filename := args[1]
		paramMap := readParams(args[2:])
		result := executeLocal(filename, paramMap)
		fmt.Printf("ExecutionResult: %q", result)
		return true
	}

	return false
}

func readParams(args []string) map[string]string {
	fmt.Printf("Args: %q", args)

	if len(args) > 0 {
		paramMap := make(map[string]string)

		for _, param := range args {

			paramContent := strings.Split(param, "=")

			if len(paramContent) == 2 {
				paramMap[paramContent[0]] = paramContent[1]
			}
		}

		return paramMap
	}

	return nil
}

func executeLocal(filename string, functionParams map[string]string) string {

	dat, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Failed loading file %q. Error: %q", filename, err)
		return ""
	}

	result, err := executor.ExecuteFunction("0", string(dat), functionParams)

	if err != nil {
		fmt.Printf("Error occured while executing function in file %q. Error: %q", filename, err)
		return ""
	}

	return result
}
