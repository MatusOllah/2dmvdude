package cache

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newPathCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "path",
		Short: "Print the path to the 2DMV cache directory",
		Long:  "Print the path to the local cache directory where 2DMVs are stored.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getCacheDir())
		},
	}
}
