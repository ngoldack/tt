package frame

import (
	"github.com/spf13/cobra"
)

var FrameCmd = &cobra.Command{
	Use:   "frame",
	Short: "prints the last 5 frames",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactive()
	},
}

func interactive() {

}

func init() {

}
