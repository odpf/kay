package main

import (
	"fmt"
	"os"

	"github.com/odpf/kay/cli"
)

func main() {
	if err := cli.New().Execute(); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
