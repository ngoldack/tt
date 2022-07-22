package cmd

import (
	"github.com/ngoldack/tt/cmd/frame"
	"github.com/ngoldack/tt/cmd/project"
	"github.com/ngoldack/tt/cmd/tag"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tt",
	Short: "timetracker root command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(frame.FrameCmd)
	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(tag.TagCmd)
}

func init() {
	addSubCommands()
}
