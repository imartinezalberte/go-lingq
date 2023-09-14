package cmd

import (
	"encoding/json"
	"reflect"

	"github.com/spf13/cobra"
)

func handleResponse(cmd *cobra.Command, response any) {
	if k := reflect.TypeOf(response).Kind(); (k == reflect.Array || k == reflect.Slice) &&
		reflect.ValueOf(response).Len() == 0 {
		cmd.Println("nothing found")
		return
	}

	buf, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	cmd.Println(string(buf))
}
