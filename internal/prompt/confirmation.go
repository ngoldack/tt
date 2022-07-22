package prompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func GetConfirmation(label string, def bool) bool {
	if def {
		label = label + " (Y/n):"
	} else {
		label = label + "(y/N):"
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate: func(input string) error {
			if len(input) > 0 {
				r := []rune(input)
				if r[0] != 'y' && r[0] != 'n' && r[0] != 'Y' && r[0] != 'N' {
					return errors.New("y or n:")
				}
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	if len(result) == 0 {
		return def
	}

	r := []rune(result)
	switch r[0] {
	case 'y':
		return true
	case 'Y':
		return true
	case 'n':
		return false
	case 'N':
		return false
	}
	return false
}
