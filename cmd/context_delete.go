/*
Copyright © 2020 Lucas Campos <lucas.campos@axoniq.io>

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
package cmd

import (
	"axon-server-cli/httpwrapper"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	contextDelete string
)

// contextDeleteCmd represents the contextDelete command
var contextDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Deletes the context.",
	Long:    `Deletes the context from axonserver.`,
	Run:     deleteContext,
}

func init() {
	contextCmd.AddCommand(contextDeleteCmd)
	contextDeleteCmd.Flags().StringVarP(&contextDelete, "context", "c", "", "*Name of the context")
	// required flags
	contextDeleteCmd.MarkFlagRequired("context")
}

func deleteContext(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("%s/v1/context/%s", viper.GetString("server"), contextDelete)
	log.Printf("calling: %s\n", url)

	responseBody := httpwrapper.DELETE(url)

	fmt.Printf("%s\n", responseBody)
}
