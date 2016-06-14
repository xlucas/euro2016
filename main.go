package main

import (
	"fmt"
	"os"

	"github.com/xlucas/euro2016/cmd"
	"github.com/xlucas/euro2016/util"
)

func main() {
	err := util.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	cmd.Initialize()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
