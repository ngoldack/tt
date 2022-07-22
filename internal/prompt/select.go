package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/ngoldack/tt/internal/db"
	"github.com/ngoldack/tt/internal/model"
	"log"
	"strings"
)

func SelectProject(label string, projects []model.Project) *model.Project {
	var items []string
	for _, p := range projects {
		items = append(items, p.Name)
	}

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range projects {
		if p.Name == result {
			return &p
		}
	}
	return nil
}

func SelectTag(label string, tags []model.Tag) *model.Tag {
	var items []string
	for _, t := range tags {
		items = append(items, t.Name)
	}

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range tags {
		if t.Name == result {
			return &t
		}
	}
	return nil
}

func SelectProjectWithActive(label string, projects []model.Project) *model.Project {
	var items []string
	for _, p := range projects {
		var item string
		if p.Active {
			item = fmt.Sprintf("%s - Active", p.Name)
		} else {
			item = fmt.Sprintf("%s - Inactive", p.Name)
		}
		items = append(items, item)
	}

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	result, _, _ = strings.Cut(result, " - Active")
	result, _, _ = strings.Cut(result, " - Inactive")

	for _, p := range projects {
		if p.Name == result {
			return &p
		}
	}
	return nil
}

func SelectTagWithAddOption(label string, tags []model.Tag) *model.Tag {
	var items []string
	for _, t := range tags {
		items = append(items, t.Name)
	}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    label,
			Items:    items,
			AddLabel: "New Tag",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, t := range tags {
		if t.Name == result {
			return &t
		}
	}
	newTag := &model.Tag{
		Name: result,
	}
	err = db.Mgr.CreateTag(newTag)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return newTag
}
