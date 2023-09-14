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
*/
package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/imartinezalberte/go-lingq/internal/language"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

var (
	languageReq language.LanguageRequest
	languageSupported bool
)

// languagesCmd represents the languages command
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "Get languages from lingq library",
	Long: `Get languages from lingq library`,
	Run: func(cmd *cobra.Command, _ []string) {
		if !cmd.Flags().Changed(language.SupportedName) {
			cmd.Println(languageReq.Supported)
			languageReq.Supported = nil
		}	else {
			*languageReq.Supported = languageSupported
		}
	
		languages, err := getLanguages(cmd)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		handleResponse(cmd, languages)
	},
}

func getLanguages(cmd *cobra.Command) (any, error) {
	client, err := rest.DefaultClient("https://www.lingq.com/api/v2")
	if err != nil {
		return nil, err
	}

	repo := language.NewRepo(client)
	service := language.NewService(repo)

	ctx, cl := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cl()

	return language.Execute(ctx, service, languageReq)
}

func init() {
	rootCmd.AddCommand(languagesCmd)

	languagesCmd.PersistentFlags().StringVarP(&languageReq.Title, language.TitleName, language.TitleShortName, "", language.TitleUsage)
	languagesCmd.PersistentFlags().StringVarP(&languageReq.Code, language.CodeName, language.CodeShortName, "", language.CodeUsage)
	languagesCmd.PersistentFlags().BoolVarP(&languageSupported, language.SupportedName, language.SupportedShortName, true, language.SupportedUsage)
	languagesCmd.PersistentFlags().IntVarP(&languageReq.ID, language.IDName, language.IDShortName, 0, language.IDUsage)
}
