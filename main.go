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
package main

import (
	"github.com/imartinezalberte/go-lingq/cmd"

	_ "github.com/imartinezalberte/go-lingq/cmd/context"
	_ "github.com/imartinezalberte/go-lingq/cmd/course"
	_ "github.com/imartinezalberte/go-lingq/cmd/course/create"
	_ "github.com/imartinezalberte/go-lingq/cmd/course/get"
	_ "github.com/imartinezalberte/go-lingq/cmd/language"
	_ "github.com/imartinezalberte/go-lingq/cmd/search"
	_ "github.com/imartinezalberte/go-lingq/cmd/search/get"
	_ "github.com/imartinezalberte/go-lingq/cmd/shelf"
	_ "github.com/imartinezalberte/go-lingq/cmd/shelf/get"
)

func main() {
	cmd.Execute()
}
