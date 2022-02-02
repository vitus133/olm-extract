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

// getCurrentCmd represents the getCurrent command
var getCurrentCmd = &cobra.Command{
	Use:   "getversions",
	Short: "Get current operator versions installed on the cluster",
	Long: `Get current operator versions installed on the cluster.
    
    Receives a list of packages as a parameter, and outputs currently running operator versions.
    Under the hood, it performs following operations:
    1. Get cluster operators, for example:
        $oc get operators -A
        NAME                                             AGE
        local-storage-operator.openshift-local-storage   4d12h
        ptp-operator.openshift-ptp                       4d12h
    2. Filter the list by the packages provided by the input parameter
    3. Get subscriptions, for example:
        $oc get sub -A
        NAMESPACE                 NAME                        PACKAGE                  SOURCE             CHANNEL
        openshift-local-storage   local-storage-operator      local-storage-operator   redhat-operators   4.9
        openshift-ptp             ptp-operator-subscription   ptp-operator             redhat-operators   4.9
    4. Get CSV version for each package of interest, for example:
        $oc get sub ptp-operator-subscription -n openshift-ptp -oyaml |grep currentCSV 
        currentCSV: ptp-operator.4.9.0-202201130525
    5. Format and print
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getCurrent called")
		packages, _ := cmd.Flags().GetStringArray("packages")

		fmt.Printf("%v\n", packages)
	},
}

func init() {
	rootCmd.AddCommand(getCurrentCmd)
	getCurrentCmd.Flags().StringArray("packages", []string{}, "A help for packages")
	getCurrentCmd.MarkFlagRequired("packages")

}
