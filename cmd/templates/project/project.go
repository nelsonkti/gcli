/**
** @创建时间 : 2022/4/14 17:24
** @作者 : fzy
 */
package project

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

type Project struct {
	Name string
	Path string
	Dest string
}

func New(name, path, dest string) *Project {
	return &Project{Name: name, Path: path, Dest: dest}
}

// 创建项目
func (p *Project) Create(repoUrl string) error {

	repo := NewRepo(repoUrl)

	if err := repo.Copy(p.Dest, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}

	return nil
}

// 判断该项目是否存在
func (p *Project) IsExists() error {
	if _, err := os.Stat(p.Dest); !os.IsNotExist(err) {
		fmt.Printf("❎️ %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}

		if !override {
			return errors.New(fmt.Sprintf("\n%s already exists. \nYou have selected the %s not to be overwritten.", p.Name, p.Name))
		}
		_ = os.RemoveAll(p.Dest)
	}

	return nil
}
