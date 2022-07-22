package project

import (
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/spf13/cobra"
	"log"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "toggles a projects state",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveToggle()
	},
}

func interactiveToggle() {
	projects, err := db.Mgr.FindProjects()
	if err != nil {
		log.Fatal(err)
	}
	p := prompt.SelectProjectWithActive("select a project", projects)
	p.Active = !p.Active
	err = db.Mgr.UpdateProject(p)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	ProjectCmd.AddCommand(toggleCmd)
}
