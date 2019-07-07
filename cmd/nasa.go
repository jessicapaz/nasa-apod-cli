/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/jessicapaz/nasa-apod-cli/api"
	"github.com/spf13/cobra"
)

// nasaCmd represents the nasa command
var nasaCmd = &cobra.Command{
	Use:   "nasa",
	Short: "gets Astronomy Picture of the Day",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data := api.GetAPOD("")
		date, _ := cmd.Flags().GetString("date")
		if date != "" {
			data = api.GetAPOD(date)
		}
		fmt.Printf("\nDate: %s\n\nTitle: %s\n\nExplanation: %s\n\nImage URL: %s\n\nCopyright: %s\n", data.Date, data.Title, data.Explanation, data.URL, data.Copyright)
		downloadImage, _ := cmd.Flags().GetBool("download-image")
		if downloadImage != false {
			imageName := "image-" + data.Date
			api.DownloadImage(data.URL, imageName)
		}
	},
}

func init() {
	rootCmd.AddCommand(nasaCmd)
	nasaCmd.Flags().StringP("date", "d", "", "Date of image")
	nasaCmd.Flags().BoolP("download-image", "i", false, "Download the APOD image")
}
