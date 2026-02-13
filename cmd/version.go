package cmd

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version and build information",
	Long:  `Display the version, git commit, build date, and runtime information for the netatmo CLI.`,
	Example: `  netatmo version
  netatmo version --short
  netatmo version --json`,
	Run: func(cmd *cobra.Command, args []string) {
		short, _ := cmd.Flags().GetBool("short")

		if short {
			fmt.Println(Version)
			return
		}

		if jsonOutput {
			info := map[string]string{
				"version":   Version,
				"commit":    Commit,
				"buildDate": BuildDate,
				"goVersion": runtime.Version(),
				"os":        runtime.GOOS,
				"arch":      runtime.GOARCH,
			}
			output, _ := json.MarshalIndent(info, "", "  ")
			fmt.Println(string(output))
			return
		}

		fmt.Printf("netatmo %s\n", Version)
		fmt.Printf("  Commit:     %s\n", Commit)
		fmt.Printf("  Built:      %s\n", BuildDate)
		fmt.Printf("  Go version: %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolP("short", "s", false, "Print only the version number")
}
