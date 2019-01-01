package main

import (
	"fmt"

	tf "testfolder"
)

func main() {
	fmt.Println("--- Running test.go")
	TestCall()
	tf.FolderPage()
	fmt.Println("--- test.go executed!")
}
