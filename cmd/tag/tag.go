package tag

import (
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"log"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "displays all tags",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := db.Mgr.FindTags()
		if err != nil {
			log.Fatal(err)
			return
		}
		if len(tags) == 0 {
			fmt.Println(util.ErrorColor("no tags found"))
			return
		}
		prompt.PrintTableTags(tags)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
