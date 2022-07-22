package tag

import (
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "removes a tag",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveRm()
	},
}

func interactiveRm() {
	tags, err := db.Mgr.FindTags()
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(tags) == 0 {
		fmt.Println(util.ErrorColor("no tags found"))
		return
	}

	t := prompt.SelectTag("select a tag", tags)
	err = db.Mgr.DeleteTag(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tag %s deleted\n", util.TagColor(t.Name))
}

func init() {
	TagCmd.AddCommand(rmCmd)
}
