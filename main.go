package main

import (
	"fmt"
	"os"

	"github.com/taiki1288/go-cobra-example/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println("%+v\n", err)
		os.Exit(1)
	}
}
