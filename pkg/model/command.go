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

// func Init(etc *Resource) *Configuration {
// 	c := &Command{
// 		Arguments: etc.Arguments,
// 		P:      tmpl.GetTemplateProvider(etc.Provider),
// 		Before: process.Command(etc.Before),
// 		After:  process.Command(etc.After),
// 	}

// 	for _, action := range etc.Actions {
// 		c.Actions = append(c.Actions, &Action{
// 			Before:   process.Command(action.Before),
// 			After:    process.Command(action.After),
// 			Arguments:   action.Arguments,
// 			P:        tmpl.GetTemplateProvider(action.Provider),
// 			To:       action.To,
// 			Template: action.Template,
// 			Copy:     action.Copy,
// 		})
// 	}

// 	return c
// }

// func (cmd *Command) Exec() (err error) {
// 	if cmd.Arguments == nil {
// 		cmd.Arguments = rtinfo.Arguments()
// 	} else {
// 		for key, val := range rtinfo.Arguments() {
// 			if _, ok := cmd.Arguments[key]; !ok {
// 				cmd.Arguments[key] = val
// 			}
// 		}
// 	}

// 	if cmd.Before != nil {
// 		if err = cmd.Before(); err != nil {
// 			return
// 		}
// 	}

// 	for _, action := range cmd.Actions {
// 		rtstatus.ExitIfError(action.Exec(cmd.Path, cmd.P, cmd.Arguments))
// 	}

// 	if cmd.After != nil {
// 		return cmd.After()
// 	}

// 	return nil
// }
