package cmd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var TheVersion string

func getVersion() string {
	if TheVersion != "" {
		return TheVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	return bi.Main.Version
}

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of 2DMVdude",
		Long: `All software has versions. This prints 2DMVdude's version.
Even Rui checks his before debugging show chaos.`,
		Run: func(cmd *cobra.Command, args []string) {
			cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
			bold := color.New(color.Bold).SprintFunc()

			ver := getVersion()
			if ver == "" {
				fmt.Fprintln(os.Stderr, "2DMVdude version is unavailable.")
			} else {
				fmt.Fprintln(os.Stderr, cyan("2DMVdude"), bold("version"), ver)
			}
			fmt.Fprintln(os.Stderr, cyan("Go"), bold("version"), runtime.Version(), "("+runtime.GOOS+"/"+runtime.GOARCH+")")
		},
	}
}

func init() {
	rootCmd.AddCommand(NewVersionCommand())
}
