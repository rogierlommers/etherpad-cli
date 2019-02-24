package cmd

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list al pads",
	Run:   listPads,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listPads(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	pads, err := etherClient.ListAllPads(ctx)
	if err != nil {
		fmt.Printf("etherpad error: %s\n", err)
		os.Exit(0)
	}

	s := reflect.ValueOf(pads.Data["padIDs"])

	if s.Len() == 0 {
		fmt.Println("you don't have any pads yet")
		os.Exit(0)
	}

	fmt.Println("Your pads:")
	for i := 0; i < s.Len(); i++ {
		stringValue := fmt.Sprintf("%s", s.Index(i))
		fmt.Printf("- %s (%s)\n", stringValue, getPadURL(stringValue))
	}

}
