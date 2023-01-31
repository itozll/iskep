package model

// Action 命令
type Action struct {
	// 执行脚本前命令列表
	Before []string `json:"before,omitempty" yaml:"before"`

	// 执行脚本后命令列表
	After []string `json:"after,omitempty" yaml:"after"`

	// 模板提供商
	Provider string `json:"provider,omitempty" yaml:"provider"`

	// // 当前脚本参数列表
	// Arguments map[string]interface{} `json:"arguments,omitempty" yaml:"arguments"`

	// 目标文件
	To string `json:"to,omitempty" yaml:"to"`
	// 源模板文件
	Template []string `json:"template,omitempty" yaml:"template"`
	// 源文件
	Copy []string `json:"copy,omitempty" yaml:"copy"`

	// // 子命令列表
	// Actions []*Action `json:"actions,omitempty" yaml:"actions"`

	Intercept bool `json:"intercept,omitempty" yaml:"intercept"`
}

// run sed ...
type Script struct {
	File  string
	Flag  string
	After bool
}
