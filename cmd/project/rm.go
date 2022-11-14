package project

import (
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/spf13/cobra"
	"log"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "removes a project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveRm()
	},
}

func interactiveRm() {
	projects, err := db.Mgr.FindProjects()
	if err != nil {
		log.Fatal(err)
	}
	p := prompt.SelectProject("Select a project", projects)
	err = db.Mgr.DeleteProject(p)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	ProjectCmd.AddCommand(rmCmd)
}
