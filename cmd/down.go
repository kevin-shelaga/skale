/*Package cmd licensing

Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/kevin-shelaga/skale/helpers"
	"github.com/kevin-shelaga/skale/k8s"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Dynamically scale all deployments down",
	Long: `Dynamically scale all deployments down to zero 
replicas. For example:

skale down`,
	Run: func(cmd *cobra.Command, args []string) {
		if helpers.IsDryRun(args) {
			fmt.Println("Dry run! No changes will be made!")
		}
		fmt.Println("Scaling down...")

		namespaces := helpers.GetNamespaces(args)

		var k k8s.KubernetesAPI = k8s.KubernetesAPI{Client: nil}

		k.Client = k.Connect()

		for _, n := range namespaces {
			deploys := k.GetDeployments(n)
			k.ScaleDeployments(deploys, nil, k8s.ScaleDown, helpers.IsDryRun(args))
		}

		fmt.Println("Done!")

	},
	DisableFlagParsing: true,
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
