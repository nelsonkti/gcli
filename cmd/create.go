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

	},
}

func CreateProject(args []string) error {
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
			return nil
		}
	} else {
		projectName = args[0]
	}

	fmt.Println(xprintf.Green("️⛹ ...\n"))

	pro := project.New(path.Base(projectName), projectName, path.Join(wd, path.Base(projectName)))

	// 判断框架是否存在
	err = pro.IsExists()
	if err != nil {
		return err
	}

	err = pro.Create(repoUrl)
	if err != nil {
		return err
	}

	fmt.Println(xprintf.Blue("project created successfully\n"))

	fmt.Println(xprintf.Green("💻 Use the following command to start the project 👇: \n"))
	fmt.Println(xprintf.Green(fmt.Sprintf("$ cd %s", pro.Dest)))
	fmt.Println(xprintf.Green("$ go mod tidy"))

	return err
}
