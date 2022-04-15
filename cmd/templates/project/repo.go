/**
** @创建时间 : 2022/4/14 17:41
** @作者 : fzy
 */
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

// 新仓库
func NewRepo(url string) *Repo {
	repo := &Repo{
		url:  url,
		home: homeDir("repo/" + repoDir(url)),
	}

	repo.source = repo.getSource()

	return repo
}

// 复制仓库
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

// 从git下载文件
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

// 从github 更新
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

// 获取文件路径
func (r *Repo) getSource() string {
	start := strings.LastIndex(r.url, "/")
	end := strings.LastIndex(r.url, ".git")

	if end == -1 {
		end = len(r.url)
	}

	return path.Join(r.home, r.url[start+1:end])
}

// 创建项目
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

// 临时存放位置
func homeDir(dir string) string {
	home := path.Join(HomeDir(), dir)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

// 仓库目录
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

// ModulePath returns go module path.
func ModulePath(filename string) (string, error) {
	modBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}

// 创建文件
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
