package main

import (
	//"github.com/onrik/logrus/filename"
	//"github.com/onrik/logrus/sentry"
	log "github.com/sirupsen/logrus"
	"github.com/john-deng/cli-demo/pkg/cmds"
)

//var (
//	dsn = "http://60a0257d7b5a429a8838e5f2ba873ec9:cb785a64cd3649ea987a1f2f5fad5e82@example.com/1"
//)

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)

	//filenameHook := filename.NewHook()
	//filenameHook.Field = "file" // Customize source field name
	//log.AddHook(filenameHook)
	//
	//sentryHook := sentry.NewHook(dsn, log.PanicLevel, log.FatalLevel, log.ErrorLevel)
	//log.AddHook(sentryHook)
}

func main() {
	cmds.RootCmd.Execute()
}
