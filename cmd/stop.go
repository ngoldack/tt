package cmd

import (
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		active, err := db.Mgr.FindFrameByActive(true)
		if err != nil {
			log.Fatal(err)
			return
		}
		if active == nil {
			fmt.Println("No active frame!")
			return
		} else {
			active.StopTime = time.Now()
			active.Active = false
			err := db.Mgr.UpdateFrame(active)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("%v stopped at %s\n", util.ProjectColor(active.Project.Name), util.TimeColor(util.FormatTime(active.StopTime)))
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
