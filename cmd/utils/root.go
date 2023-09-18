package utils

import (
	"encoding/json"
	"reflect"

	"github.com/spf13/cobra"
)

func HandleResponse(cmd *cobra.Command, response any, err error) {
	if response == nil {
		cmd.PrintErrln("no response")
		return
	}

	if err != nil {
		cmd.PrintErrln(err)
		return
	}

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
