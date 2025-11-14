package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/MatusOllah/2dmvdude/internal/adbutil"
	"github.com/MatusOllah/2dmvdude/internal/mv"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func pull(id int, dst string, serial string, kind mv.MVKind, region mv.ServerRegion, fallbackKind bool, force bool) error {
	// open device
	if verbose {
		verbosePrintf("Opening device at ADB %s with serial number %s...\n", adbAddress, strconv.Quote(serial))
	}
	device, err := adbutil.OpenDevice(adbAddress, serial)
	if err != nil {
		return fmt.Errorf("opening device: %v\nTry starting the ADB server using this command:\n\tadb start-server", err)
	}

	if verbose {
		verbosePrintln("Device opened successfully.")
		adbutil.PrintDeviceInfo(device)
	}

	// get path
	path := mv.MVPath(id, kind, region)

	// get file size
	size, err := adbutil.GetRemoteFileSize(device, path)
	if err != nil {
		if fallbackKind {
			altKind := kind.Other()

			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Fprintf(os.Stderr, "⚠️  %s %s %s %s%s\n", yellow("Requested 2DMV kind"), kind.String(), yellow("not found, trying fallback kind"), altKind, yellow("..."))

			altPath := mv.MVPath(id, altKind, region)

			size, err = adbutil.GetRemoteFileSize(device, altPath)
			if err != nil {
				return fmt.Errorf("2DMV of either kind not found, failed to stat file: %w", err)
			}

			path = altPath
			//kind = altKind
		} else {
			return fmt.Errorf("2DMV not found, failed to stat file: %w", err)
		}
	}

	// create output file
	if dst == "" {
		dst = fmt.Sprintf("%04d.usm", id)
	}

	if _, err := os.Stat(dst); err == nil {
		if force {
			yellow := color.New(color.FgYellow).SprintFunc()
			// The emoji renders weirdly, hence the 2 spaces
			fmt.Fprintf(os.Stderr, "⚠️  %s %s %s\n", yellow("File"), dst, yellow("already exists and will be overwritten."))
		} else {
			return fmt.Errorf("file %s already exists. Use --force to overwrite it", dst)
		}
	}

	file, err := os.Create(dst)
	checkErr(err)
	defer func() {
		if err := file.Close(); err != nil {
			checkErr(err)
		}
	}()

	if verbose {
		verbosePrintf("Pulling remote file %s => %s (%d bytes)...\n", path, dst, size)
	}

	// create progress bar
	pb := progressbar.DefaultBytes(size, "Pulling CRI Sofdec video file...")

	// pull the video
	if err := device.Pull(path, io.MultiWriter(file, pb)); err != nil {
		return fmt.Errorf("failed to pull file %s: %w", path, err)
	}

	return nil
}

func NewPullCommand() *cobra.Command {
	var (
		serial       string
		kind         mv.MVKind = mv.MVKindSEKAI
		fallbackKind bool
		region       mv.ServerRegion = mv.ServerRegionEN
		output       string
		force        bool
	)

	c := &cobra.Command{
		Use:     "pull id",
		Short:   "Pull a raw 2DMV from an Android device",
		Long:    `Pull a raw CRI Sofdec 2DMV from an Android device using the specified song ID.`,
		Example: `2dmvdude pull 264`, // MORE MORE JUMP! - Parasol Cider (#264)
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parse song ID integer
			id, err := strconv.Atoi(args[0])
			if err != nil {
				checkErr(fmt.Errorf("failed to parse song ID %s: %w", strconv.Quote(args[0]), err))
			}

			if id <= 0 {
				checkErr(fmt.Errorf("song ID must be a positive integer"))
			}

			checkErr(pull(id, output, serial, kind, region, fallbackKind, force))

			// success message
			fmt.Fprintln(os.Stderr, color.GreenString("✅ Successfully pulled 2DMV! -> %04d", id))
		},
	}

	// Flags
	c.Flags().StringVarP(&serial, "serial", "s", "", "Device serial number")
	c.Flags().VarP(&kind, "kind", "k", "Type of 2DMV to prefer pulling (\"original\", \"sekai\")")
	c.Flags().BoolVar(&fallbackKind, "fallback", true, "Fallback to other kind if requested kind not found")
	c.Flags().VarP(&region, "region", "r", "Game server region (\"jp\", \"en\", \"tw\", \"kr\", \"cn\")")
	c.Flags().StringVarP(&output, "output", "o", "", "Output file path (default <id>.usm)")
	c.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite existing output file")

	return c
}

func init() {
	rootCmd.AddCommand(NewPullCommand())
}
