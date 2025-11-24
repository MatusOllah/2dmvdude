package cache

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func newWipeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "wipe",
		Short: "Wipe the 2DMV cache",
		Long:  "Wipe (delete) all files in the local 2DMV cache directory.",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(os.RemoveAll(getCacheDir()))
		},
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, color.HiRedString("Error:"), err)
		os.Exit(1)
	}
}
