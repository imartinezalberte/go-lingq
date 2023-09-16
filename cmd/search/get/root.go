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
	cmd "github.com/imartinezalberte/go-lingq/cmd/search"
	"github.com/spf13/cobra"
)

// GetSearchCmd represents the search command
var GetSearchCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
}

func init() {
	cmd.SearchCmd.AddCommand(GetSearchCmd)
}
