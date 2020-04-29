// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"MiniHttpFileServer/pkg/img"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload image and return markdown style link",
	Long: `Upload your images to website https://sm.ms and return 
a Markdown style link. It upload .jpg .pug et.`,
	Run: func(cmd *cobra.Command, args []string) {
		i := &img.Img{
			Path: P,
		}
		b, err := i.Upload()
		if err != nil {
			log.Error(err)
		} else {
			if len(b) < 200 {
				fmt.Print(string(b))
			}
			url, err := img.GetUrl(b)
			if err != nil {
				log.Error(err)
				os.Exit(124)
			}
			fmt.Println(fmt.Sprintf("![%s](%s)", url, url))
		}
	},
}
var P string

func init() {
	uploadCmd.Flags().StringVarP(&P, "path", "p", "", "Path of image")
	uploadCmd.MarkFlagRequired("path")
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
