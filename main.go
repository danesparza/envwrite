package main

import (
	"flag"
	"log"
	"os"
)

//	Set up our flags
var (
	environmentVariable = flag.String("env", "PATH", "The environment variable to read")
	filePath            = flag.String("file", "c:\\temp\\file.txt", "The full path of the file to write to")
)

func main() {

	//	Parse the command line for flags:
	flag.Parse()

	//	Prep the file to write to:
	f, err := os.Create(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//	Write the environment variable to the file:
	f.WriteString(os.Getenv(*environmentVariable))

}
