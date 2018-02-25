package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/john-deng/cli-demo/pkg/cmds"
)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)

}

func main() {
	cmds.RootCmd.Execute()
}
