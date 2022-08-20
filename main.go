package main

import (
	"fmt"
	"os"

	"github.com/odpf/kay/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
