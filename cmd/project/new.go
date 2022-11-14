package project

import (
	"errors"
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/model"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "creates a new project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveCr()
	},
}

func interactiveCr() {
	project := &model.Project{Active: true}

	project.Name = prompt.GetInput("project name:", func(input string) error {
		if len(input) == 0 {
			return errors.New("please provide a name")
		}
		_, err := db.Mgr.FindProjectByName(input)
		if err != nil {
			return err
		}
		return nil
	})

	project.TicketNr = prompt.GetInput("ticket-nr:", func(input string) error {
		if len(input) == 0 {
			return errors.New("please provide a name")
		}
		return nil
	})

	project.Comment = prompt.GetInput("comment (optional):", func(input string) error {

		return nil
	})

	if prompt.GetConfirmation("Are you sure?", true) {
		if err := db.Mgr.CreateProject(project); err != nil {
			log.Fatal("an error occurred while creating project", err)
		}
	}

	fmt.Printf("%s created!\n", util.ProjectColor(project.Name))
}

func init() {
	ProjectCmd.AddCommand(newCmd)
}
