package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Application closed !!!")
	os.Exit(0)
	return nil
}
