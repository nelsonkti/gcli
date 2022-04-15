/**
** @创建时间 : 2022/4/14 17:24
** @作者 : fzy
 */
package project

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"path"
)

type Project struct {
	Name string
	Path string
}

func New(name string, path string) *Project {
	return &Project{Name: name, Path: path}
}

// 创建项目
func (p *Project) Create(dir, repoUrl string) error {
	to := path.Join(dir, p.Name)

	// 判断框架是否存在
	err := p.isExists(to)
	if err != nil {
		return err
	}

	repo := NewRepo(repoUrl)

	if err := repo.Copy(to, p.Path, []string{".git", ".github"}); err != nil {
		return err
	}

	return nil
}

// 判断该项目是否存在
func (p *Project) isExists(to string) error {
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
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
			return err
		}
		_ = os.RemoveAll(to)
	}

	return nil
}
