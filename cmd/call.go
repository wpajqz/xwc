package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wpajqz/xwc/config"
)

// RunCallCommand 调用命令
func RunCallCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "call",
		Short:   "Call command",
		Aliases: []string{"c"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.DisableFlagParsing = true
			cmd.DisableFlagsInUseLine = true
			c := &config.Config{}
			if err := c.LoadConfigFile("xwc.yml"); err != nil {
				return err
			}

			for _, element := range c.Enviroment {
				for k, v := range element {
					if err := os.Setenv(k, v); err != nil {
						return err
					}
				}
			}

			if len(args) > 0 {
				command := args[0]
				if v, ok := c.Command[command]; ok {
					var param []string

					ss := strings.Split(v, " ")
					if len(ss) > 1 {
						param = append(param, ss[1:]...)
					}

					if len(args) > 1 {
						param = append(param, args[1:]...)
					}

					fmt.Printf("exec: %s %s\n", ss[0], strings.Join(param, " "))

					var out bytes.Buffer
					execCmd := exec.Command(ss[0], param...)
					execCmd.Stdout = &out
					execCmd.Stderr = &out

					err := execCmd.Run()
					if err != nil {
						fmt.Println(out.String())
						return err
					}

					fmt.Println(out.String())

					return nil
				}

				return fmt.Errorf("cmd %s isn't exists", command)
			}

			return errors.New("cmd is empty")
		},
	}
}
