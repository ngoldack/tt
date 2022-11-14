package project

import (
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/spf13/cobra"
	"log"
)

var ProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "prints all projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := db.Mgr.FindProjects()
		if err != nil {
			log.Fatal(err)
		}
		if len(projects) != 0 {
			prompt.PrintTableProjects(projects)
		}
	},
}

func init() {
}
