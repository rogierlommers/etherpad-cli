package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a pad",
	Run:   deletePad,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deletePad(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	if len(args) == 0 {
		fmt.Println("please specify pad name")
		os.Exit(0)
	}

	resp, err := etherClient.DeletePad(ctx, args[0])
	if err != nil {
		fmt.Printf("etherpad error: %s\n", err)
		os.Exit(0)
	}

	fmt.Println(resp.Message)
}
