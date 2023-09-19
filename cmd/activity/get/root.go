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
	"github.com/imartinezalberte/go-lingq/cmd/activity"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	a "github.com/imartinezalberte/go-lingq/internal/activity"
	"github.com/imartinezalberte/go-lingq/internal/config"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/spf13/cobra"
)

var activityReq ActivityRequest

// getActivityCmd represents the activity command
var getActivityCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, _ []string) {
		activities, err := getActivities()
		utils.HandleResponse(cmd, activities, err)
	},
}

func getActivities() (any, error) {
	client, err := rest.DefaultClient(config.BaseURLV2)
	if err != nil {
		return nil, err
	}

	repo := a.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := a.NewService(repo)

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	return a.Execute(ctx, service, activityReq)
}

func init() {
	activity.ActivityCmd.AddCommand(getActivityCmd)
	activityReq.Args(getActivityCmd)
}
