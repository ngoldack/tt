package project

import (
	"errors"
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edits a project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveEdit()
	},
}

func interactiveEdit() {
	projects, err := db.Mgr.FindProjects()
	if err != nil {
		log.Fatal(err)
	}
	p := prompt.SelectProject("select a project", projects)

	p.Name = prompt.GetInputWithDefault("project Name:", p.Name, func(input string) error {
		if len(input) == 0 {
			return errors.New("please provide a name")
		}
		_, err := db.Mgr.FindProjectByName(input)
		if err != nil {
			return err
		}
		return nil
	})

	p.TicketNr = prompt.GetInputWithDefault("ticket-nr:", p.TicketNr, func(input string) error {
		if len(input) == 0 {
			return errors.New("please provide a name")
		}
		return nil
	})

	p.Comment = prompt.GetInputWithDefault("comment (Optional):", p.Comment, func(input string) error {
		return nil
	})

	if prompt.GetConfirmation("are you sure?", true) {
		if err := db.Mgr.UpdateProject(p); err != nil {
			log.Fatal("an error occurred while creating project", err)
			return
		}
	}

	fmt.Printf("%s edited successfully\n", util.ProjectColor(p.Name))
}

func init() {
	ProjectCmd.AddCommand(editCmd)
}
