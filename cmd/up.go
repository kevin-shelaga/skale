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

	"github.com/kevin-shelaga/skale/k8s"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Dynamically scale all deployments up",
	Long: `Dynamically scale all deployments up to the 
minimum replicas. For example:

skale up`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Scaling up...")
		client := k8s.Connect()
		deploys := k8s.GetDeployments(client)
		hpas := k8s.GetHorizontalPodAutoscalers(client)
		k8s.ScaleDeployments(client, deploys, hpas, k8s.ScaleUp)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
