package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/wpajqz/xwc/config"
)

// RunInitCommand 初始化项目命令
func RunInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Init config file",
		Aliases: []string{"i"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c := &config.Config{
				Enviroment: make([]map[string]string, 0),
				Command:    make(map[string]string),
			}

			c.Enviroment = append(c.Enviroment, map[string]string{"your_enviroment": "your_enviroment"})
			c.Command = map[string]string{"your_command": "your_command"}

			if c.IsExists("xwc.yml") {
				return errors.New("xwc.yml is exists")
			}

			return c.StoreConfigFile("xwc.yml")
		},
	}
}
