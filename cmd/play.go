package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func isLocalFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func NewPlayCommand() *cobra.Command {
	var (
		ffplayArgs string
	)

	c := &cobra.Command{
		Use:   "play",
		Short: "Play a 2DMV",
		Long:  `Play a 2DMV either from a local file or from an Android device by specifying the song ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			isFile := isLocalFile(args[0])

			if isFile {
				extraArgs := []string{}
				if ffplayArgs != "" {
					extraArgs = strings.Fields(ffplayArgs)
				}
				playFile(args[0], extraArgs)
			}
		},
	}

	// Flags
	c.Flags().StringVarP(&ffplayArgs, "ffplay-args", "a", "", "Additional arguments to pass to FFplay")

	return c
}

func playFile(path string, extraArgs []string) {
	// Play local file
	cmd := exec.Command("ffplay", append([]string{
		"-i", path,
		"-autoexit",
		"-hide_banner",
		"-loglevel", "error",
		"-stats",
		"-volume", "100",
		"-window_title", filepath.Base(path) + " - 2DMVdude",
	}, extraArgs...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if verbose {
		verbosePrintln("FFplay command:", cmd.String())
	}

	fmt.Fprintln(os.Stderr, "Playing...")
	checkErr(cmd.Run())
}

func init() {
	rootCmd.AddCommand(NewPlayCommand())
}
