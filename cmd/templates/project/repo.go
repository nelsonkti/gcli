// Package project
// @Author fzy
// @Date 2022-04-14 15:04:03
package project

import (
	"bytes"
	"fmt"
	"github.com/nelsonkti/gcli/util/xstring"
	"golang.org/x/mod/modfile"
	"log"
	stdurl "net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Repo struct {
	url    string
	home   string
	source string
	dest   string
}

// NewRepo
// @Description: 新仓库
// @param url
// @return *Repo
func NewRepo(url string) *Repo {
	repo := &Repo{
		url:  url,
		home: homeDir("repo/" + repoDir(url)),
	}

	repo.source = repo.getSource()

	return repo
}

// Copy
// @Description: 复制仓库
// @receiver r
// @param dest
// @param modPath
// @param ignores
// @return error
func (r *Repo) Copy(dest string, modPath string, ignores []string) error {

	if err := r.clone(); err != nil {
		return err
	}

	modulePath, err := ModulePath(path.Join(r.source, "go.mod"))

	if err != nil {
		return err
	}

	replaces := []string{modulePath, modPath}

	r.dest = dest

	err = r.createProject(r.source, dest, replaces, ignores)
	if err != nil {
		return err
	}

	if err = os.RemoveAll(HomeDir()); err != nil {
		return err
	}

	return err
}

// clone
// @Description: 从git下载文件
// @receiver r
// @return error
func (r *Repo) clone() error {
	if _, err := os.Stat(r.source); !os.IsNotExist(err) {
		return r.pull()
	}

	cmd := exec.Command("git", "clone", r.url, r.source)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

// pull
// @Description: 从github 更新
// @receiver r
// @return error
func (r *Repo) pull() error {

	cmd := exec.Command("git", "symbolic-ref", "HEAD")
	cmd.Path = r.source
	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}
	cmd = exec.Command("git", "pull")
	cmd.Dir = r.source
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return err

}

// getSource
// @Description: 获取文件路径
// @receiver r
// @return string
func (r *Repo) getSource() string {
	start := strings.LastIndex(r.url, "/")
	end := strings.LastIndex(r.url, ".git")

	if end == -1 {
		end = len(r.url)
	}

	return path.Join(r.home, r.url[start+1:end])
}

// createProject
// @Description: 创建项目
// @receiver r
// @param source
// @param dest
// @param replaces
// @param ignores
// @return error
func (r *Repo) createProject(source string, dest string, replaces []string, ignores []string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(source); err != nil {
		fmt.Println(err)
		return err
	}

	if err := os.MkdirAll(dest, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(source); err != nil {
		return err
	}

	// 遍历目录和文件
	for _, value := range fds {
		// 忽略一些文件
		if xstring.InArray(value.Name(), ignores) {
			continue
		}

		sourceSrc := path.Join(source, value.Name())
		destSrc := path.Join(dest, value.Name())

		if value.IsDir() {
			if err := r.createProject(sourceSrc, destSrc, replaces, ignores); err != nil {
				return err
			}
		} else {
			if err := createFile(sourceSrc, destSrc, replaces); err != nil {
				return err
			}
		}
	}

	return err
}

// homeDir
// @Description: 临时存放位置
// @param dir
// @return string
func homeDir(dir string) string {
	home := path.Join(HomeDir(), dir)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

// repoDir
// @Description: 仓库目录
// @param url
// @return string
func repoDir(url string) string {
	if !strings.Contains(url, "//") {
		url = "//" + url
	}
	if strings.HasPrefix(url, "//git@") {
		url = "ssh:" + url
	} else if strings.HasPrefix(url, "//") {
		url = "https:" + url
	}
	u, err := stdurl.Parse(url)
	if err == nil {
		url = fmt.Sprintf("%s://%s%s", u.Scheme, u.Hostname(), u.Path)
	}
	var start int
	start = strings.Index(url, "//")
	if start == -1 {
		start = strings.Index(url, ":") + 1
	} else {
		start += 2
	}
	end := strings.LastIndex(url, "/")
	return url[start:end]
}

// ModulePath
// @Description: 返回go module path
// @param filename
// @return string
// @return error
func ModulePath(filename string) (string, error) {
	modBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}

// createFile
// @Description: 创建文件
// @param src
// @param dst
// @param replaces
// @return error
func createFile(src, dst string, replaces []string) error {
	var err error
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	buf, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	var old string
	for i, next := range replaces {
		if i%2 == 0 {
			old = next
			continue
		}
		buf = bytes.ReplaceAll(buf, []byte(old), []byte(next))
	}

	return os.WriteFile(dst, buf, srcinfo.Mode())
}
