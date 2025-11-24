package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/MatusOllah/2dmvdude/internal/mv"
	"github.com/spf13/cobra"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func NewPlayCommand() *cobra.Command {
	var (
		ffplayArgs   string
		serial       string
		kind         mv.MVKind = mv.MVKindSEKAI
		fallbackKind bool
		region       mv.ServerRegion = mv.ServerRegionEN
		skipLeadin   bool
	)

	c := &cobra.Command{
		Use:     "play",
		Short:   "Play a 2DMV",
		Long:    `Play a 2DMV either from a local file or from an Android device by specifying the song ID.`,
		Example: `2dmvdude play 264`, // MORE MORE JUMP! - Parasol Cider (#264)
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("you must specify either a file path or a song ID")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			extraArgs := strings.Fields(ffplayArgs)

			cachePath := filepath.Join(os.TempDir(), "2dmvdude/mvcache", args[0]+".usm")

			if fileExists(args[0]) {
				playFile(args[0], skipLeadin, extraArgs)
			} else if fileExists(cachePath) {
				playFile(cachePath, skipLeadin, extraArgs)
			} else {
				// Fetch and play remote file

				// parse song ID integer
				id, err := strconv.Atoi(args[0])
				if err != nil {
					checkErr(fmt.Errorf("failed to parse song ID %s: %w", strconv.Quote(args[0]), err))
				}

				if id <= 0 {
					checkErr(fmt.Errorf("song ID must be a positive integer"))
				}

				checkErr(pull(id, cachePath, serial, kind, region, fallbackKind, true))

				playFile(cachePath, skipLeadin, extraArgs)
			}
		},
	}

	// Flags
	c.Flags().StringVarP(&ffplayArgs, "ffplay-args", "a", "", "Additional arguments to pass to FFplay")
	c.Flags().StringVarP(&serial, "serial", "s", "", "Device serial number")
	c.Flags().VarP(&kind, "kind", "k", "Type of 2DMV to prefer pulling (\"original\", \"sekai\")")
	c.Flags().BoolVar(&fallbackKind, "fallback", true, "Fallback to other kind if requested kind not found")
	c.Flags().VarP(&region, "region", "r", "Game server region (\"jp\", \"en\", \"tw\", \"kr\", \"cn\")")
	c.Flags().BoolVar(&skipLeadin, "skip-leadin", true, "Skip first 9 seconds of silence")

	return c
}

func playFile(path string, skipLeadin bool, extraArgs []string) {
	// Play local file
	args := []string{
		"-i", path,
		"-autoexit",
		"-hide_banner",
		"-loglevel", "error",
		"-stats",
		"-volume", "100",
		"-window_title", filepath.Base(path) + " - 2DMVdude",
	}
	if skipLeadin {
		args = append(args, "-ss", "9")
	}
	cmd := exec.Command("ffplay", append(args, extraArgs...)...)
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
