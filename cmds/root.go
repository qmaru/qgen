package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	projectName string
	rootCmd     = &cobra.Command{
		Use:     "qgen",
		Short:   "Generate code",
		Version: "20230703",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "Project Name")
	rootCmd.MarkFlagRequired("name")
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(
		CreateCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
