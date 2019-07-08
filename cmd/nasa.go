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
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		dImg, _ := cmd.Flags().GetBool("download-image")
		if start != "" && end != "" {
			data := api.GetAPODs(api.DateRange{Start: start, End: end})
			for _, value := range data {
				printFmtData(value)
				if dImg != false {
					downloadImage(value)
				}
			}
		} else {
			date, _ := cmd.Flags().GetString("date")
			data := api.GetAPOD(date)
			printFmtData(data)
			if dImg != false {
				downloadImage(data)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(nasaCmd)
	nasaCmd.Flags().StringP("date", "d", "", "Image's date")
	nasaCmd.Flags().StringP("start", "s", "", "Image's start date")
	nasaCmd.Flags().StringP("end", "e", "", "Image's end date")
	nasaCmd.Flags().BoolP("download-image", "i", false, "Download the APOD image")
}

func printFmtData(value api.ResponseData) {
	fmt.Printf("\nDate: %s\n\nTitle: %s\n\nExplanation: %s\n\nImage URL: %s\n\nCopyright: %s\n", value.Date, value.Title, value.Explanation, value.URL, value.Copyright)
}

func downloadImage(value api.ResponseData) {
	imageName := "image-" + value.Date
	api.DownloadImage(value.URL, imageName)
}
