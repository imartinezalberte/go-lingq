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
*/package language

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/imartinezalberte/go-lingq/cmd"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	"github.com/imartinezalberte/go-lingq/internal/config"
	"github.com/imartinezalberte/go-lingq/internal/language"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

var (
	languageReq       LanguageRequest
	languageSupported bool
)

// getLanguagesCmd represents the languages command
var getLanguagesCmd = &cobra.Command{
	Use:   "get",
	Short: "Get languages from lingq library",
	Long:  `Get languages from lingq library`,
	Run: func(cmd *cobra.Command, _ []string) {
		if !cmd.Flags().Changed(SupportedName) {
			cmd.Println(languageReq.Supported)
			languageReq.Supported = nil
		} else {
			*languageReq.Supported = languageSupported
		}

		languages, err := getLanguages()
		utils.HandleResponse(cmd, languages, err)
	},
}

func getLanguages() (any, error) {
	client, err := rest.DefaultClient(config.BaseURL)
	if err != nil {
		return nil, err
	}

	repo := language.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := language.NewService(repo)

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	return language.Execute(ctx, service, languageReq)
}

func init() {
	cmd.RootCmd.AddCommand(languagesCmd)

	Args(getLanguagesCmd, &languageReq, &languageSupported)
}

func Args(cmd *cobra.Command, target *LanguageRequest, languageSupported *bool) {
	cmd.Flags().
		StringVarP(&target.Title, TitleName, TitleShortName, TitleDefault, TitleUsage)
	cmd.Flags().
		StringVarP(&target.Code, CodeName, CodeShortName, CodeDefault, CodeUsage)
	cmd.Flags().
		BoolVarP(languageSupported, SupportedName, SupportedShortName, SupportedDefault, SupportedUsage)
	cmd.Flags().
		UintVarP(&target.ID, IDName, IDShortName, IDDefault, IDUsage)
	cmd.Flags().
		UintVarP(&target.KnownWords, KnownWordsName, KnownWordsShortName, KnownWordsDefault, KnownWordsUsage)
}
