package actions

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A CLI task manager",
}
