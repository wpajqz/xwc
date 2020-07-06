package main

import (
	"bytes"
	"fmt"
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
	if err := c.LoadConfigFile("xwc.yml"); err != nil {
		fmt.Println(err.Error())
		return
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
				var param []string

				ss := strings.Split(v.(string), " ")
				if len(ss) > 1 {
					param = append(param, ss[1:]...)
				}

				param = append(param, args...)

				for k, v := range param {
					param[k] = os.ExpandEnv(v)
				}

				fmt.Printf("exec: %s %s\n", ss[0], strings.Join(param, " "))

				var ob, eb bytes.Buffer
				execCmd := exec.Command(ss[0], param...)
				execCmd.Stdout = &ob
				execCmd.Stderr = &eb

				err := execCmd.Run()
				if err != nil {
					fmt.Println(eb.String())
					return
				}

				fmt.Println(ob.String())
			},
		}

		subCmd.DisableFlagParsing = true
		rootCmd.AddCommand(subCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
