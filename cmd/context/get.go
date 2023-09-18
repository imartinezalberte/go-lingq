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
package context

import (
	"time"

	"github.com/spf13/cobra"

	contxt "context"

	"github.com/imartinezalberte/go-lingq/cmd"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	"github.com/imartinezalberte/go-lingq/internal/config"
	con "github.com/imartinezalberte/go-lingq/internal/context"
	"github.com/imartinezalberte/go-lingq/internal/rest"
)

var contextReq ContextRequest

// getContextCmd represents the context command
var getContextCmd = &cobra.Command{
	Use:   "get",
	Short: "Helps to retrieve the possible contexts which user can use",
	Long:  `Helps to retrieve the possible contexts which user can use`,
	Run: func(cmd *cobra.Command, _ []string) {
		contexts, err := getContexts()
		utils.HandleResponse(cmd, contexts, err)
	},
}

func getContexts() (any, error) {
	client, err := rest.DefaultClient(config.BaseURLV2)
	if err != nil {
		return nil, err
	}

	repo := con.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := con.NewService(repo)

	ctx, cl := contxt.WithTimeout(contxt.Background(), 10*time.Second)
	defer cl()

	return con.Execute(ctx, service, contextReq)
}

func init() {
	contextCmd.AddCommand(getContextCmd)

	contextReq.Args(getContextCmd)
}
