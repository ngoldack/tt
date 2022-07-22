package tag

import (
	"errors"
	"fmt"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/model"
	"github.com/ngoldack/tt/internal/prompt"
	"github.com/ngoldack/tt/internal/util"
	"github.com/spf13/cobra"
	"strings"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "creates a new tag",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			interactiveNew()
		} else if len(args) > 1 {
			fmt.Println(util.ErrorColor("too many arguments"))
		}
		argumentNew(args[0])
	},
}

func argumentNew(name string) {
	if len(name) == 0 {
		fmt.Println(util.ErrorColor("name must be specified"))
		return
	}
	t := &model.Tag{
		Name: name,
	}
	err := db.Mgr.CreateTag(t)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: tags.name") {
			fmt.Println(util.ErrorColor("tag with this name already exists"))
			return
		}
		fmt.Println(util.ErrorColor("error creating tag"))
		return
	}
	fmt.Printf("%s created successfully\n", util.TagColor(name))
}

func interactiveNew() {
	name := prompt.GetInput("Name:", func(input string) error {
		if len(input) == 0 {
			return errors.New("provide a name")
		}

		if t, err := db.Mgr.FindTagByName(input); err == nil {
			if t != nil {
				return errors.New("tag with this name already exists")
			}
		}

		return nil
	})

	ok := prompt.GetConfirmation(fmt.Sprintf("creating %s", util.TagColor(name)), true)
	if !ok {
		return
	}

	t := &model.Tag{
		Name: name,
	}
	err := db.Mgr.CreateTag(t)
	if err != nil {
		fmt.Println(util.TagColor("error creating tag"))
		return
	}
}

func init() {
	TagCmd.AddCommand(newCmd)
}
