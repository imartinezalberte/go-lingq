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
package get

import (
	"context"
	"time"

	"github.com/imartinezalberte/go-lingq/cmd"
	searchCmd "github.com/imartinezalberte/go-lingq/cmd/search"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	"github.com/imartinezalberte/go-lingq/internal/config"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/imartinezalberte/go-lingq/internal/search"
	"github.com/spf13/cobra"
)

var searchResources SearchResources

// GetSearchCmd represents the search command
var GetSearchCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		resources, err := getResources()
		utils.HandleResponse(cmd, resources, err)
	},
}

func getResources() (any, error) {
	client, err := rest.DefaultClient(config.BaseURL)
	if err != nil {
		return nil, err
	}

	repo := search.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := search.NewService(repo)

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	return search.Execute(ctx, service, searchResources)
}

func init() {
	searchCmd.SearchCmd.AddCommand(GetSearchCmd)
	searchResources.Args(GetSearchCmd)
}
