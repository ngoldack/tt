package cmd

import (
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the current status of the frame",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		active, err := db.Mgr.FindFrameByActive(true)
		if err != nil {
			return
		}
		if active == nil {
			fmt.Println("No active frame")
			return
		}

		fmt.Printf("%s with %s is active since %s.\n", util.ProjectColor(active.Project.Name), util.TagColor(active.Tag.Name), util.TimeColor(util.FormatTime(active.StartTime)))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
