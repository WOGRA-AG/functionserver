package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("functionmanager-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	err = initWebService()

	if err != nil {
		log.Fatal(err)
	}
}
