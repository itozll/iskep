package options

import (
	"github.com/itozll/iskep/runtime/iflag"
)

var (
	GroupName    = iflag.NewString(nil, "group", "g", "", "group name")
	GoVersion    = iflag.NewString(nil, "go-version", "", "1.17", "go version")
	Verbose      = iflag.NewBool(nil, "verbose", "V", false, "add more details to output logging")
	DryRun       = iflag.NewBool(nil, "dry-run", "", false, "run through and reports activity without writing out results")
	File         = iflag.NewString(nil, "file", "f", "", "customize action with file")
	FileType     = iflag.NewString(nil, "file-type", "", "yaml", "file type, support json/yaml")
	Local        = iflag.NewBool(nil, "local", "", false, "initialize git repository on current directory")
	Arg          = iflag.NewStringToString(nil, "arg", "a", nil, "specify variable list")
	ExpansionArg = iflag.NewStringToString(nil, "arg-expansion", "e", nil, "pecify extended variable list")
)
