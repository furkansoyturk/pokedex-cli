package main

import (
	"fmt"
	"os"
)
func commandExit() error  {
	fmt.Println("Application closed !!!")
	os.Exit(0)
	return nil
}