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
package course

import (
	"github.com/imartinezalberte/go-lingq/cmd"
	"github.com/spf13/cobra"
)

// CourseCmd represents the course command
var CourseCmd = &cobra.Command{
	Use:   "course",
	Short: "Handle CRUD of courses on lingq",
	Long:  `Handle CRUD of courses on lingq`,
}

func init() {
	cmd.RootCmd.AddCommand(CourseCmd)
}
