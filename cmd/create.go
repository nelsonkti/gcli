package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/nelsonkti/gcli/cmd/templates/project"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	repoUrl = ""
)

func init() {

	Cmd.AddCommand(createProject)

	createProject.PersistentFlags().StringVar(&repoUrl, "url", "https://github.com/nelsonkti/echo-framework.git", "go code library")

}

var createProject = &cobra.Command{
	Use:   "create",
	Short: "create a new project",
	Long: `create a new project command line: gcli create demo 
modify the project path: gcli create demo --path projectPath`,
	Run: func(cmd *cobra.Command, args []string) {

		err := CreateProject(args)

		if err != nil {
			fmt.Println(xprintf.Red(err.Error()))
			return
		}

		fmt.Println(xprintf.Blue("project created successfully"))

	},
}

func CreateProject(args []string) (err error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var projectName string
	if len(args) == 0 {
		prompt := &survey.Input{
			Message: "What is project name ?",
			Help:    "Created project name. ",
		}
		err = survey.AskOne(prompt, &projectName)
		if err != nil || projectName == "" {
			return
		}
	} else {
		projectName = args[0]
	}

	project := project.New(path.Base(projectName), projectName)

	project.Create(wd, repoUrl)

	return nil
}
