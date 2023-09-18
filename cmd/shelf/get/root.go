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
*/package get

import (
	"context"
	"time"

	"github.com/imartinezalberte/go-lingq/cmd"
	shelfCmd "github.com/imartinezalberte/go-lingq/cmd/shelf"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	"github.com/imartinezalberte/go-lingq/internal/config"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/imartinezalberte/go-lingq/internal/shelf"
	"github.com/spf13/cobra"
)

var shelfRequest ShelfRequest

// getShelvesCmd represents the shelf get command
var getShelvesCmd = &cobra.Command{
	Use:   "get",
	Short: "Getting shelves from lingq",
	Long:  `Gettign shelves from lingq`,
	Run: func(cmd *cobra.Command, _ []string) {
		shelfRes, err := getShelves()
		utils.HandleResponse(cmd, shelfRes, err)
	},
}

func getShelves() (any, error) {
	client, err := rest.DefaultClient(config.BaseURL)
	if err != nil {
		return nil, err
	}

	repo := shelf.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := shelf.NewService(repo)

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	return shelf.Execute(ctx, service, shelfRequest)
}

func init() {
	shelfCmd.ShelfCmd.AddCommand(getShelvesCmd)
	shelfRequest.Args(getShelvesCmd)
}
