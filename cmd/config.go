package cmd

import (
	"github.com/spf13/viper"
	"github.com/xlucas/euro2016/util"
)

var (
	client    *util.JSONClient
	showEmoji bool
)

func Initialize() {
	client = util.NewJSONClient(endpoint, viper.GetString("token"))
	showEmoji = viper.GetBool("emoji")
}
