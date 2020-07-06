package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/elliotchance/orderedmap"
	"github.com/spf13/cobra"
	"github.com/wpajqz/xwc/cmd"
	"github.com/wpajqz/xwc/config"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "xwc",
		Short: "wrap some command with enviroment",
	}

	rootCmd.AddCommand(cmd.RunInitCommand())

	c := &config.Config{}
	if c.IsExists("xwc.yml") {
		if err := c.LoadConfigFile("xwc.yml"); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for _, element := range c.Enviroment {
		for k, v := range element {
			if err := os.Setenv(k, v); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	om := orderedmap.NewOrderedMap()
	for k, v := range c.Command {
		om.Set(k, v)
	}

	for _, k := range om.Keys() {
		v, _ := om.Get(k)

		subCmd := &cobra.Command{
			Use:   k.(string),
			Short: fmt.Sprintf("Exec: %s", v.(string)),
			Long:  fmt.Sprintf("Exec: %s", v.(string)),
			Run: func(cmd *cobra.Command, args []string) {
				command, param := parseFlag(v.(string), args)

				fmt.Printf("exec: %s %s\n", command, strings.Join(param, " "))

				execCmd := exec.Command(command, param...)

				stdout, err := execCmd.StdoutPipe()
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				defer stdout.Close()

				err = execCmd.Start()
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				reader := bufio.NewReader(stdout)

				for {
					line, err := reader.ReadString('\n')
					if err != nil || err == io.EOF {
						break
					}

					fmt.Println(line)
				}

				if err := execCmd.Wait(); err != nil {
					fmt.Println(err.Error())
				}

				return
			},
		}

		subCmd.DisableFlagParsing = true
		rootCmd.AddCommand(subCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}

// parseFlag 解析命令以及命令所对应的参数
func parseFlag(flag string, args []string) (command string, param []string) {
	ss := strings.Split(flag, " ")
	if len(ss) > 1 {
		param = append(param, ss[1:]...)
	}

	param = append(param, args...)

	for k, v := range param {
		param[k] = os.ExpandEnv(v)
	}

	return ss[0], param
}
