package prompt

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ngoldack/tt/internal/model"
)

var (
	headerProject = table.Row{"ID", "Active", "Name", "Ticket-Nr", "Comment"}
	headerTag     = table.Row{"ID", "Name"}
)

func PrintTableProjects(projects []model.Project) {
	t := table.NewWriter()
	t.AppendHeader(headerProject)

	for _, project := range projects {
		t.AppendRow(table.Row{project.ID, project.Active, project.Name, project.TicketNr, project.Comment})
	}

	fmt.Println(t.Render())
}

func PrintTableTags(tags []model.Tag) {
	t := table.NewWriter()
	t.AppendHeader(headerTag)

	for _, tag := range tags {
		t.AppendRow(table.Row{tag.ID, tag.Name})
	}

	fmt.Println(t.Render())
}
