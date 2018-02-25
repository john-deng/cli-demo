package cmds

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"os"
)

var optionFoo string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run command of cli-demo",
	Long:  `the run command of cli-demo`,
	Run:   handleRunCmd,
}

func init() {
	runCmd.PersistentFlags().StringVar(&optionFoo, "foo", "", "option foo")
	RootCmd.AddCommand(runCmd)
}

func handleRunCmd(cmd *cobra.Command, args []string)  {
	log.Info("cli-demo run ...")
	wd, err := os.Getwd()
	if nil == err {
		log.Debugf("working dir: %s", wd)
	}

	// for demo only
	if "" != optionFoo {
		log.Infof("option foo: %s", optionFoo)
	}
}