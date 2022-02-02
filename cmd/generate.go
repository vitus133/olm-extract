/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the images pull list",
	Long: `Generates the list of image pull specifications based on required 
    packages + channels and optionally current package versions. 
    If current packages versions are provided, the latest upgradable version
    is selected from the channel metadata. If not, latest version in the channel is selected.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringArray("packages-and-channels", []string{},
		`A comma-separated list of required packages and channels, for example:
    "ptp-operator:stable,sriov-network-operator:stable"`)
	generateCmd.MarkFlagRequired("packages-and-channels")
	generateCmd.Flags().StringArray("indexes", []string{},
		`A comma-separated list of indexes, for example:
        "quay.io/example-org/example-catalog:v1,registry.redhat.io/redhat/redhat-operator-index:v4.9"`)
	generateCmd.MarkFlagRequired("indexes")
	generateCmd.Flags().StringArray("csvs", []string{},
		`A comma-separated list current csv versions, for example:
        "ptp-operator.4.9.0-202201130525,local-storage-operator.4.9.0-202112142229"`)
}
