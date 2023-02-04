package rtinfo

import (
	"fmt"
	"strings"

	"github.com/itozll/iskep/pkg/runtime/rtstatus"
)

var Info = struct {
	GoVersion string

	// Workspace = [Group/]<Project>
	Workspace string

	// <Domain>/<Group>/<Project>
	Repository string

	// default github.com
	Domain string
	Group  string

	// 项目名称
	Project string

	// 项目目录
	Directory string
}{
	Domain: "github.com",
}

var (
	SkipGit    bool
	File       string
	FileType   string
	Parent     string
	TargetPath string
)

func Binder() map[string]string {
	return map[string]string{
		"go_version": Info.GoVersion,
		"workspace":  Info.Workspace,
		"repository": Info.Repository,
		"domain":     Info.Domain,
		"group":      Info.Group,
		"project":    Info.Project,
		"directory":  Info.Directory,
	}
}

func Init(workspace string) error {
	group, project := split(workspace)
	if project == "" {
		return fmt.Errorf("nil workspace name")
	}

	if Info.Group == "" {
		Info.Group = group
		if Info.Group == "" {
			return fmt.Errorf("empty group name")
		}
	}

	Info.Project = project
	Info.Directory = project
	Info.Workspace = Info.Group + "/" + Info.Project
	Info.Repository = Info.Domain + "/" + Info.Workspace

	return nil
}

func split(reposName string) (string, string) {
	list := strings.Split(reposName, "/")
	if len(list) > 2 {
		rtstatus.Fatal("error repos_name: %s", reposName)
	}

	if len(list) == 1 {
		return "", list[0]
	}

	return list[0], list[1]
}
