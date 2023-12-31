/*
Copyright © 2023 imartinezalberte

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	EnvPrefix string = "LINGQ"

	// LINGQ_TOKEN
	TokenEnv string = "TOKEN"
)

var Token string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-lingq",
	Short: "lingq go tool to handle connections with lingq api",
	Long:  `lingq go tool to handle connections with lingq api`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix(EnvPrefix)
	viper.AutomaticEnv()

	RootCmd.PersistentFlags().
		StringVar(&Token, "token", viper.GetString(TokenEnv), "That's a long string that you can find at https://www.lingq.com/en/accounts/apikey/ on your own account")

	RootCmd.MarkFlagRequired("token")
}
