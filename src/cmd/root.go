// Copyright S 2022 Jean-Francois Gratton <jean-francois@famillegratton.net>

package cmd

import (
	"filenormalizer/executor"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands

var (
	version      = "1.000 (2023.08.03), J.F.Gratton <jean-francois@famillegratton.net>"
	uppercase    = false
	lowercase    = false
	verboseMode  = false
	normalize    = false
	stripPattern = []string{}
)

var rootCmd = &cobra.Command{
	Use:     "filenormalizer [directory]",
	Version: version,
	Short:   "Batch file renaming tool to lowercase",
	Long: `Rename files in named directory in parameter.
If no directory is specified, the current directory will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		var targets = []string{}
		if len(args) == 0 {
			targets[0] = "."
		} else {
			targets = args
		}
		executor.Rename(verboseMode, normalize, uppercase, lowercase, stripPattern, targets)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolVarP(&verboseMode, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&uppercase, "uppercase", "u", false, "rename file to uppercase")
	rootCmd.PersistentFlags().BoolVarP(&lowercase, "lowercase", "l", false, "rename file to lowercase")
	rootCmd.PersistentFlags().BoolVarP(&lowercase, "normalize", "n", false, "normalize filenames")
	rootCmd.PersistentFlags().StringSliceVarP(&stripPattern, "strip", "s", []string{}, "Strip pattern from file names")
}
