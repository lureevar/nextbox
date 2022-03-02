package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var conf struct {
	Path string
}

var rootCmd = &cobra.Command{
	Use:   "nextbox",
	Short: "A simple todo list app for the CLI",
	Long:  `A simple todo list application for the CLI based on the CRUD principle`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	dir := os.Getenv("XDG_DATA_HOME")
	if len(dir) == 0 {
		dir = os.Getenv("HOME")
	}

	full := dir + "/" + rootCmd.Name() + `.csv`

	viper.AddConfigPath("$XDG_CONFIG_HOME")
	viper.AddConfigPath("$HOME")

	viper.SetConfigName(".nextbox")
	viper.SetConfigType("toml")

	viper.SetDefault("path", full)

	viper.SafeWriteConfig()
	viper.Unmarshal(&conf)
}
