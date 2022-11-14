package cmd

import (
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/model"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveStart()
	},
}

func interactiveStart() {
	f, err := db.Mgr.FindFrameByActive(true)
	if err != nil {
		log.Fatal(err)
		return
	}
	if f != nil {
		fmt.Println(util.ErrorColor("a frame is currently active"))
		return
	}
	projects, err := db.Mgr.FindProjectsByActive(true)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(projects) == 0 {
		fmt.Println(util.ErrorColor("no active projects found"))
		return
	}
	p := prompt.SelectProject("select a project", projects)

	tags, err := db.Mgr.FindTags()
	t := prompt.SelectTagWithAddOption("select a tag or create a new one:", tags)
	if err != nil {
		log.Fatal(err)
		return
	}

	f = &model.Frame{
		Active:    true,
		StartTime: time.Now(),
		Project:   p,
		Tag:       t,
		Comment:   "",
	}

	err = db.Mgr.CreateFrame(f)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("starting %s with %s at %s\n", util.ProjectColor(p.Name), util.TagColor(f.Tag.Name), util.TimeColor(util.FormatTime(f.StartTime)))
}

func init() {
	rootCmd.AddCommand(startCmd)
}
