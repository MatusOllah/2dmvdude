package cache

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCacheCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "cache",
		Short: "Manage the 2DMV cache",
		Long: `Manage the local cache of 2DMVs.
This command allows you to wipe the cache or view its status.`,
	}

	c.AddCommand(newPathCommand())

	return c
}

func getCacheDir() string {
	return filepath.Join(os.TempDir(), "2dmvdude", "mvcache")
}
