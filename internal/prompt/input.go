package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func GetInput(label string, validateFunc func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validateFunc,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

func GetInputWithDefault(label string, def string, validateFunc func(input string) error) string {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: templates,
		Validate:  validateFunc,
		Default:   def,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
