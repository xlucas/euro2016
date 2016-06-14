package main

import (
	"fmt"
	"os"

	"github.com/xlucas/euro2016/cmd"
	"github.com/xlucas/euro2016/util"
)

func main() {
	util.LoadConfig()
	cmd.Initialize()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
