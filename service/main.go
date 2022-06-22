package main

import (
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 && argsWithoutProg[0] != "-service" {
		runAsLocalExecutor(argsWithoutProg)
	} else {
		runAsService()
	}
}
