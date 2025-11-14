package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func checkTool(name, cmd string) error {
	fmt.Fprintf(os.Stderr, "Checking %s... ", name)

	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "not found\n")
		return fmt.Errorf("tool %s not found in PATH: %w", name, err)
	}

	fmt.Fprint(os.Stderr, color.HiGreenString("OK"))
	if verbose {
		color.New(color.Faint).Fprintf(os.Stderr, " (%s)", path)
	}
	fmt.Fprintln(os.Stderr)

	return nil
}

func NewDoctorCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Check whether the dependencies are installed",
		Long: `Check whether the dependencies are installed and configured correctly.
More specifically, this checks for the presence of ADB, FFmpeg, and FFplay.`,
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(checkTool("ADB", "adb"))
			checkErr(checkTool("FFmpeg", "ffmpeg"))
			checkErr(checkTool("FFplay", "ffplay"))
		},
	}
}

func init() {
	rootCmd.AddCommand(NewDoctorCommand())
}
