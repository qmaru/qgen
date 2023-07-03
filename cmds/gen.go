package cmds

import (
	"fmt"
	"log"

	"github.com/qmaru/qgen/services"

	"github.com/spf13/cobra"
)

var (
	CreateCmd = &cobra.Command{
		Use:     "create",
		Short:   "Generate code",
		Version: "20230703",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	webCmd = &cobra.Command{
		Use:   "web",
		Short: "Generate web",
		Run: func(cmd *cobra.Command, args []string) {
			project := services.Project{
				Name: projectName,
				Type: cmd.Use,
			}
			err := project.Gen()
			if err != nil {
				log.Fatal(err)
			}
			tips(projectName, cmd.Use)
		},
	}
	cliCmd = &cobra.Command{
		Use:   "cli",
		Short: "Generate cli",
		Run: func(cmd *cobra.Command, args []string) {
			project := services.Project{
				Name: projectName,
				Type: cmd.Use,
			}
			err := project.Gen()
			if err != nil {
				log.Fatal(err)
			}
			tips(projectName, cmd.Use)
		},
	}
)

func init() {
	CreateCmd.AddCommand(webCmd)
	CreateCmd.AddCommand(cliCmd)
}

func tips(pname, ptype string) {
	fmt.Printf("Create project with %s: %s\n", ptype, pname)
	fmt.Printf("cd %s\n", pname)
	fmt.Println("go mod init")
	fmt.Println("go mod tidy")
}
