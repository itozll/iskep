package options

import "github.com/itozll/iskep/pkg/iflag"

var (
	GroupName = iflag.NewString(nil, "group", "g", "", "group name")
	GoVersion = iflag.NewString(nil, "go-version", "", "1.17", "go version")
	Verbose   = iflag.NewBool(nil, "verbose", "V", false, "add more details to output logging")
	DryRun    = iflag.NewBool(nil, "dry-run", "", false, "run through and reports activity without writing out results")
)
