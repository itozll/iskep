package model

type Command struct {
	Name        string `json:"name,omitempty" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description"`
	Style       string `json:"style,omitempty" yaml:"style"`

	// 执行脚本前命令列表
	Before []string `json:"before,omitempty" yaml:"before"`

	// 执行脚本后命令列表
	After []string `json:"after,omitempty" yaml:"after"`

	Provider string `json:"provider,omitempty" yaml:"provider"`

	Arguments []*Argument `json:"arguments,omitempty" yaml:"arguments"`

	Actions []*Action `json:"actions,omitempty" yaml:"actions"`

	Intercept string `json:"intercept,omitempty" yaml:"intercept"`
}
