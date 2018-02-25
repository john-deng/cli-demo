package cmds

import (
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use:     "cli-demo",
	Short:   "Command line program for cli-demo",
	Long:    "Command line program for cli-demo",
	Example: `  $ cli-demo foo`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Usage()

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-demo/config.yml)")
	//RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "John Deng <john.deng@utlook.com>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli-demo/config")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	appName := viper.Get("app.name")
	log.Debugf("app.name: %s", appName)
	appAuthor := viper.Get("app.author")
	log.Debugf("app.author: %s", appAuthor)
}
