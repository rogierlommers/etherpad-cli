package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new pad",
	Run:   AddPad,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func AddPad(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	if len(args) == 0 {
		fmt.Println("please specify pad name")
		os.Exit(0)
	}

	resp, err := etherClient.CreatePad(ctx, args[0], "")
	if err != nil {
		fmt.Printf("etherpad error: %s", err)
		os.Exit(0)
	}

	fmt.Printf("%s, link: %s\n", resp.Message, getPadURL(args[0]))
}
