package model

type Argument struct {
	Name        string   `json:"name,omitempty" yaml:"name"`
	Value       string   `json:"value,omitempty" yaml:"value"`
	Description string   `json:"description,omitempty" yaml:"description"`
	Range       []string `json:"range,omitempty" yaml:"range"`

	Required  bool `json:"required,omitempty" yaml:"required"`
	Disabled  bool `json:"disabled,omitempty" yaml:"disabled"`
	Expansion bool `json:"expansion,omitempty" yaml:"expansion"`
	Fixed     bool `json:"fixed,omitempty" yaml:"fixed"`

	changed bool
}

func (arg *Argument) Change()       { arg.changed = true }
func (arg *Argument) Changed() bool { return arg.changed }
