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
	contxt "context"
	"time"

	"github.com/imartinezalberte/go-lingq/cmd"
	"github.com/imartinezalberte/go-lingq/cmd/course"
	"github.com/imartinezalberte/go-lingq/cmd/utils"
	"github.com/imartinezalberte/go-lingq/internal/config"
	cour "github.com/imartinezalberte/go-lingq/internal/course"
	"github.com/imartinezalberte/go-lingq/internal/rest"
	"github.com/spf13/cobra"
)

var coursesReq CoursesRequest

// getCoursesCmd represents the course command
var getCoursesCmd = &cobra.Command{
	Use:   "get",
	Short: "Handle creation of courses on lingq",
	Long:  `Handle creation of courses on lingq`,
	Run: func(cmd *cobra.Command, _ []string) {
		courseRes, err := getCourses()
		utils.HandleResponse(cmd, courseRes, err)
	},
}

func getCourses() (any, error) {
	client, err := rest.DefaultClient(config.BaseURL)
	if err != nil {
		return nil, err
	}

	repo := cour.NewRepo(client.SetHeader("Authorization", "Token "+cmd.Token))
	service := cour.NewService(repo)

	ctx, cl := contxt.WithTimeout(contxt.Background(), 10*time.Second)
	defer cl()

	return cour.Execute(ctx, service, coursesReq)
}

func init() {
	course.CourseCmd.AddCommand(getCoursesCmd)

	coursesReq.Args(getCoursesCmd)
}
