package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "write mandatory settings in config file",
	Run:   writeConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func writeConfig(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	// get etherpad server hostname
	fmt.Println("Enter baseURL of etherpad (without API endpoint). for example: https://myetherpad.server.com")
	fmt.Print("hostname: ")
	etherHostname, _ := reader.ReadString('\n')
	etherHostname = strings.Replace(etherHostname, "\n", "", -1)

	apiPath, err := getAPIPath(etherHostname)
	if err != nil {
		fmt.Printf("it seems like you provided the wrong hostname: %s\n", err)
		os.Exit(0)
	}

	// get etherpad token
	fmt.Println("Enter etherpad API token.")
	fmt.Print("token: ")
	etherToken, _ := reader.ReadString('\n')
	etherToken = strings.Replace(etherToken, "\n", "", -1)

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configFile := path.Join(home, fmt.Sprintf("%s.json", configFile))

	newConfig := viper.New()
	newConfig.AddConfigPath(configFile)
	newConfig.Set("etherpad_hostname", apiPath)
	newConfig.Set("etherpad_token", etherToken)

	if err := newConfig.WriteConfigAs(configFile); err != nil {
		fmt.Printf("error writing to config file (%s): %s\n", configFile, err)
		os.Exit(0)
	}

	fmt.Printf("config file %s written\n", configFile)
}

func getAPIPath(h string) (string, error) {
	apiPath := fmt.Sprintf("%s/api", h)
	fmt.Printf("api path: %s\n", apiPath)

	resp, err := http.Get(apiPath)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("non 200 response code received from %q", h)
	}

	return apiPath, nil
}
