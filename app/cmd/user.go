package cmd

import (
	"fmt"
	"log"

	"github.com/linuxsuren/jenkins-cli/client"
	"github.com/spf13/cobra"
)

type UserOption struct {
	OutputOption
}

var userOption UserOption

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringVarP(&userOption.Format, "output", "o", "json", "Format the output")
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Print the user of your Jenkins",
	Long:  `Print the user of your Jenkins`,
	Run: func(cmd *cobra.Command, args []string) {
		jenkins := getCurrentJenkins()
		jclient := &client.UserClient{}
		jclient.URL = jenkins.URL
		jclient.UserName = jenkins.UserName
		jclient.Token = jenkins.Token
		jclient.Proxy = jenkins.Proxy
		jclient.ProxyAuth = jenkins.ProxyAuth

		var tokenName string
		if len(args) > 0 {
			tokenName = args[0]
		}

		if status, err := jclient.Create(tokenName); err == nil {
			var data []byte
			if data, err = userOption.Output(status); err == nil {
				fmt.Printf("%s\n", string(data))
			} else {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	},
}
