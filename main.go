package main

import (
	"fmt"
	"os"

	"github.com/xlucas/euro2016/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
