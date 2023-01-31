package options

import (
	"github.com/itozll/iskep/pkg/runtime/iflag"
	"github.com/itozll/iskep/pkg/runtime/rtinfo"
)

var (
	GroupName = iflag.NewString(&rtinfo.Info.Group, "group", "g", "", "group name")
	GoVersion = iflag.NewString(&rtinfo.Info.GoVersion, "go-version", "", "1.19", "go version")
	Verbose   = iflag.NewBool(nil, "verbose", "V", false, "add more details to output logging")
	DryRun    = iflag.NewBool(nil, "dry-run", "", false, "run through and reports activity without writing out results")
	File      = iflag.NewString(&rtinfo.File, "file", "f", "", "customize action with file")
	FileType  = iflag.NewString(&rtinfo.FileType, "file-type", "", "yaml", "file type, support json/yaml")
	Local     = iflag.NewBool(nil, "local", "", false, "initialize git repository on current directory")
	Arg       = iflag.NewStringToString(nil, "arg", "a", nil, "specify variable list")
	Parent    = iflag.NewString(&rtinfo.Parent, "parent", "", "root", "parent command")
	Path      = iflag.NewString(&rtinfo.TargetPath, "path", "p", "", "current directory")
)
