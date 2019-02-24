package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/FabianWe/etherpadlite-golang"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const configFile = ".etherpad-cli"

var etherClient *etherpadlite.EtherpadLite

var rootCmd = &cobra.Command{
	Use:   "etherpad-cli",
	Short: "List, add and delete pads on your etherpad instance",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if len(os.Args) != 1 && os.Args[1] != "setconfig" {
		cobra.OnInitialize(initConfig)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".etherpad-cli" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(configFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file (%s): %s\n", viper.ConfigFileUsed(), err)
		os.Exit(0)
	}

	// setup etherpad client
	etherClient = etherpadlite.NewEtherpadLite(viper.GetString("etherpad_token"))
	etherClient.BaseURL = viper.GetString("etherpad_hostname")
}

func getPadURL(name string) string {

	u, err := url.Parse(etherClient.BaseURL)
	if err != nil {
		fmt.Printf("error occured: %s", err)
		os.Exit(0)
	}

	return fmt.Sprintf("%s://%s/p/%s", u.Scheme, u.Host, name)
}
