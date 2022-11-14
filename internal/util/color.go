package util

import "github.com/fatih/color"

var (
	ProjectColor = color.New(color.FgBlue, color.Bold).SprintFunc()
	TagColor     = color.New(color.FgGreen, color.Bold).SprintFunc()
	TimeColor    = color.New(color.FgYellow, color.Bold).SprintFunc()
	ErrorColor   = color.New(color.FgRed, color.Bold).SprintfFunc()
)
