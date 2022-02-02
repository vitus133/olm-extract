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
	"strings"

	"github.com/operator-framework/operator-registry/alpha/model"
	"github.com/spf13/cobra"
	idxo "github.com/vitus133/olm-extract/pkg"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the images pull list",
	Long: `
    Generates a list of image pull specifications based on the required 
    packages + channels and optionally current package versions. 
    If current packages versions are provided, the latest upgradable version
    is selected from the channel metadata. If not, latest version in the channel is selected.`,
	Run: func(cmd *cobra.Command, args []string) {
		packagesAndChannels, err := cmd.Flags().GetStringArray("packages-and-channels")
		packages := make(map[string]string)
		var filter []string
		for _, pkg := range packagesAndChannels {
			pc := strings.Split(pkg, ":")
			packages[pc[0]] = pc[1]
			filter = append(filter, pc[0])
		}
		cobra.CheckErr(err)
		indexes, err := cmd.Flags().GetStringArray("indexes")
		cobra.CheckErr(err)
		// Result is *declcfg.DeclarativeConfig concatenated from all
		// indexes and filtered by packages nd channels provided
		var model model.Model
		model, err = idxo.ProcessRefs(indexes, filter)
		cobra.CheckErr(err)
		for p, c := range packages {
			fmt.Println(">>>>>>>>>")
			currentPkg := model[p]
			fmt.Printf("%+v\n", currentPkg.Channels[c].Name)
			fmt.Printf("%+v\n", currentPkg.Channels[c].Package.Channels["stable"].Bundles)
			fmt.Printf("%+v\n", currentPkg.Channels[c].Bundles)

		}
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
