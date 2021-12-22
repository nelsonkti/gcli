package cmd

import (
	"errors"
	"fmt"
	"github.com/nelsonkti/gcli/lib"
	"github.com/nelsonkti/gcli/util/xfile"
	"github.com/nelsonkti/gcli/util/xprintf"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	Name      string
	Path      string
	ShortPath string
}

var (
	project            Project
	defaultProjectName = "project-demo"
	framework          = ""
)

func init() {

	Cmd.AddCommand(createProject)

	createProject.PersistentFlags().StringVar(&project.Path, "path", "$GOPATH", "create project path")

	createProject.PersistentFlags().StringVar(&framework, "fw", "echo", "")

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

		fmt.Println(xprintf.Blue("Project dir:", project.Path))
		fmt.Println(xprintf.Blue("project created successfully"))

	},
}

// 创建项目
func CreateProject(newArgs []string) (err error) {

	if lib.Version() < "1.13" {
		return errors.New("the current version is relatively low")
	}

	if framework != "iris" && framework != "echo" {
		return errors.New("error selection of framework type")
	}

	if len(newArgs) <= 0 {
		fmt.Println(xprintf.Red("execution failed. entered the command incorrectly, please use gcli create -h for details"))
		return
	}

	name := newArgs[0]
	if name == "" {
		project.Name = defaultProjectName
	} else {
		project.Name = name
	}

	if project.Path != "" {

		if project.Path, err = filepath.Abs(project.Path); err != nil {
			return
		}

		project.Path = filepath.Join(project.Path, project.Name)

	} else {

		path, _ := os.Getwd()
		project.Path = filepath.Join(path, project.Name)

	}

	project.ShortPath = xfile.GetModPath(project.Path)

	if err = doCreateProject(); err != nil {
		return
	}

	return
}

func doCreateProject() (err error) {

	if err = os.MkdirAll(project.Path, 0755); err != nil {
		return
	}

	templatesPath := fmt.Sprintf("./templates/" + framework)

	box := packr.New(templatesPath, templatesPath)

	for _, name := range box.List() {

		if project.ShortPath != "" && name == "go.mod.tmpl" {
			continue
		}

		tmpl, _ := box.FindString(name)

		i := strings.LastIndex(name, string(os.PathSeparator))

		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(project.Path, dir), 0755); err != nil {
				return
			}
		}

		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}

		if err = xfile.WriteFile(filepath.Join(project.Path, name), tmpl, project); err != nil {
			return
		}
	}

	return
}
